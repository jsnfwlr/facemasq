package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jsnfwlr/facemasq/api/lib/files"
	"github.com/jsnfwlr/facemasq/api/lib/password"
	_ "github.com/mattn/go-sqlite3"
)

const DBTargetVer = 1

var (
	Conn                                      *sqlx.DB
	DBPATH, DataFile, DataRoot, adminPassword string
)

func init() {
	adminPassword = os.Getenv("ADMINPASSWORD")
	if adminPassword == "" {
		adminPassword = "ResetMe"
	}
}

func Connect(dataRoot, dbFile string) (err error) {
	var failedQuery string

	DataRoot = dataRoot
	DataFile = dbFile

	DBPATH = fmt.Sprintf("%s%s", DataRoot, DataFile)

	if _, err = os.Stat(DBPATH); err == nil {
		log.Println("Database exists, connecting")
		Conn, err = sqlx.Connect("sqlite3", DBPATH)
		if err != nil {
			return
		}

	} else if os.IsNotExist(err) {
		log.Println("Creating database")
		Conn, err = sqlx.Connect("sqlite3", DBPATH)
		if err != nil {
			return
		}
		log.Println("Connecting to database")
		failedQuery, err = prepareDB()
		if err != nil {
			err = fmt.Errorf("%q: %s", err, failedQuery)
			return
		}
	}

	err = checkDBVersion()

	return
}

func checkDBVersion() (err error) {
	var name string
	var dbVerStr string
	var dbCurrentVer int
	sql := `SELECT name FROM sqlite_master WHERE type='table' AND name='Meta';`
	err = Conn.Get(&name, sql)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return
		}
		dbCurrentVer = 0
	}
	if name == "Meta" {
		sql = `SELECT Value FROM Meta WHERE Name = 'DBVersion' AND UserID IS NULL;`
		err = Conn.Get(&dbVerStr, sql)
		if err != nil {
			return
		}
		dbCurrentVer, _ = strconv.Atoi(dbVerStr)
	}
	if (DBTargetVer - dbCurrentVer) > 0 {
		log.Printf("Current DB is %d versions behind required version (v%d -> v%d)\n", (DBTargetVer - dbCurrentVer), dbCurrentVer, DBTargetVer)
	} else {
		log.Printf("Current DB version is up to date (v%d -> v%d)\n", dbCurrentVer, DBTargetVer)
	}
	for i := 1; i <= (DBTargetVer - dbCurrentVer); i++ {
		log.Printf("Upgrading DB to Version %d\n", i)
		err = upgradeDBTo(i)
		if err != nil {
			return
		}
	}
	return

}

func upgradeDBTo(dbUpgradeVer int) (err error) {
	var fileContents []byte

	backupFileName := fmt.Sprintf("%s_ver%d.sqlite", strings.Replace(DBPATH, ".sqlite", "", -1), (dbUpgradeVer - 1))
	log.Printf("Creating backup of DB  %s\n", backupFileName)
	_, err = files.Copy(DBPATH, backupFileName)
	if err != nil {
		return
	}

	sqlFileLocation := fmt.Sprintf("%s/ver_%d.sql", strings.Replace(DataRoot, "data", "upgrades", -1), dbUpgradeVer)
	if _, err = os.Stat(sqlFileLocation); err == nil {
		fileContents, err = ioutil.ReadFile(sqlFileLocation)
		if err != nil {
			return
		}
		_, err = Conn.Exec(string(fileContents))
	} else if os.IsNotExist(err) {
		err = fmt.Errorf("could not find db upgrade file %s", sqlFileLocation)
	}
	return
}

func prepareDB() (query string, err error) {
	var newPassword string
	log.Println("Initializing database")
	newPassword, err = password.HashPassword(adminPassword)
	if err != nil {
		log.Printf("error while creating password hash: %v\n", err)
		return
	}
	create := []string{
		// ParamData  - The various tables used by foreign key relationships to define/categorise the devices/interfaces/addresses/hostnames
		`CREATE TABLE "Categories" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Label" TEXT NOT NULL UNIQUE, "Icon" TEXT NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL);`,
		`CREATE TABLE "Status" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Label" TEXT NOT NULL UNIQUE, "Icon" TEXT NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL);`,
		`CREATE TABLE "DeviceTypes" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Label" TEXT NOT NULL UNIQUE, "Icon" TEXT NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL);`,
		`CREATE TABLE "InterfaceTypes" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Label" TEXT NOT NULL UNIQUE, "Icon" TEXT NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL);`,
		`CREATE TABLE "VLANs" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Label" TEXT NOT NULL UNIQUE, "IPv4Mask" TEXT NOT NULL, "IPv6Mask" TEXT NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL);`,
		`CREATE TABLE "Locations" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Label" TEXT NOT NULL UNIQUE, "IsCloud" BOOLEAN DEFAULT FALSE NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL);`,
		`CREATE TABLE "OperatingSystems" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Vendor" TEXT NOT NULL, "Family" TEXT NOT NULL, "Version" TEXT NOT NULL, "Name" TEXT NOT NULL, "IsOpenSource" BOOLEAN DEFAULT FALSE NOT NULL, "IsServer" BOOLEAN DEFAULT FALSE NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL, UNIQUE("Vendor", "Family", "Version", "Name"));`,
		`CREATE TABLE "Architectures" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Label" TEXT NOT NULL UNIQUE, "BitSpace" INTEGER DEFAULT 64 NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL);`,

		// UserData   - Users and maintainer records: users can log in to the app, maintainers are responsible for devices
		`CREATE TABLE "Users" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Username" TEXT UNIQUE, "Password" TEXT, "Label" TEXT NOT NULL UNIQUE, "CanAuthenticate" BOOLEAN DEFAULT FALSE NOT NULL, "AccessLevel" INTEGER DEFAULT 0 NOT NULL, "IsInternal" BOOLEAN DEFAULT FALSE NOT NULL, "Notes" TEXT, "IsLocked" BOOLEAN DEFAULT FALSE NOT NULL);`,

		// MetaData   - App settings & User preferences
		`CREATE TABLE "Meta" ("Name" TEXT NOT NULL, "Value" TEXT NOT NULL, "UserID" INTEGER,  FOREIGN KEY("UserID") REFERENCES "Users"("ID"), UNIQUE("Name", "UserID"));`,

		// DeviceData - Tables that define the devices, their interfaces, addresses, and hostnames
		`CREATE TABLE "Devices" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "MachineName" TEXT DEFAULT "Unknown" NOT NULL, "Brand" TEXT, "Model" TEXT, "Purchased" DATETIME, "Serial" TEXT, "IsTracked" BOOLEAN DEFAULT FALSE NOT NULL, "FirstSeen" DATETIME NOT NULL, "IsGuest" BOOLEAN DEFAULT FALSE NOT NULL, "IsOnline" BOOLEAN DEFAULT FALSE NOT NULL, "Label" TEXT, "Notes" TEXT, "CategoryID" INTEGER DEFAULT 1 NOT NULL, "StatusID" INTEGER DEFAULT 1 NOT NULL, "MaintainerID" INTEGER DEFAULT 1 NOT NULL, "LocationID" INTEGER DEFAULT 1 NOT NULL, "DeviceTypeID" INTEGER DEFAULT 1 NOT NULL, "OperatingSystemID" INTEGER DEFAULT 1 NOT NULL, "ArchitectureID" INTEGER DEFAULT 1 NOT NULL, FOREIGN KEY("CategoryID") REFERENCES "Categories"("ID"), FOREIGN KEY("StatusID") REFERENCES "Status"("ID"), FOREIGN KEY("MaintainerID") REFERENCES "Maintainers"("ID"), FOREIGN KEY("LocationID") REFERENCES "Locations"("ID"), FOREIGN KEY("DeviceTypeID") REFERENCES "DeviceTypes"("ID"), FOREIGN KEY("OperatingSystemID") REFERENCES "OperatingSystems"("ID"), FOREIGN KEY("ArchitectureID") REFERENCES "Architectures"("ID"));`,
		`CREATE TABLE "Interfaces" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "MAC" TEXT NOT NULL UNIQUE, "IsPrimary" BOOLEAN DEFAULT TRUE NOT NULL, "IsVirtual" BOOLEAN DEFAULT FALSE NOT NULL, "IsOnline" BOOLEAN DEFAULT FALSE NOT NULL, "Label" TEXT, "Notes" TEXT, "LastSeen" DATETIME NOT NULL, "StatusID" INTEGER DEFAULT 1 NOT NULL, "InterfaceTypeID" INTEGER DEFAULT 1 NOT NULL, "VLANID" INTEGER DEFAULT 1 NOT NULL, "DeviceID" INTEGER NOT NULL, FOREIGN KEY("StatusID") REFERENCES "Status"("ID"), FOREIGN KEY("InterfaceTypeID") REFERENCES "InterfaceTypes"("ID"), FOREIGN KEY("VLANID") REFERENCES "VLANs"("ID"), FOREIGN KEY("DeviceID") REFERENCES "Devices"("ID"));`,
		`CREATE TABLE "Addresses" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "IPv4" TEXT, "IPv6" TEXT, "IsPrimary" BOOLEAN DEFAULT TRUE NOT NULL, "IsVirtual" BOOLEAN DEFAULT FALSE NOT NULL, "IsReserved" BOOLEAN DEFAULT FALSE NOT NULL, "LastSeen" DATETIME NOT NULL, "Label" TEXT, "Notes" TEXT, "InterfaceID" INTEGER NOT NULL, UNIQUE("InterfaceID", "IPv4"), FOREIGN KEY("InterfaceID") REFERENCES "Interfaces"("ID"));`,
		`CREATE TABLE "Hostnames" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Hostname" TEXT NOT NULL UNIQUE, "IsDNS" BOOLEAN DEFAULT FALSE NOT NULL, "IsSelfSet" BOOLEAN DEFAULT FALSE NOT NULL, "Notes" TEXT, "AddressID" INTEGER NOT NULL, FOREIGN KEY("AddressID") REFERENCES "Addresses"("ID"));`,

		// ScanData   - Tables that record the times that addresses were online
		`CREATE TABLE "Scans" ("ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "Time" DATETIME NOT NULL UNIQUE);`,
		`CREATE TABLE "History" ("AddressID" INTEGER NOT NULL, "ScanID" TEXT NOT NULL, UNIQUE("AddressID", "ScanID"), FOREIGN KEY("AddressID") REFERENCES "Addresses"("ID"), FOREIGN KEY("ScanID") REFERENCES "Scans"("ID"));`,
	}

	insert := []string{
		`INSERT INTO "Categories" ("Label", "Icon", "Notes", "IsLocked") VALUES ("Unsorted", "HelpCircle", NULL, 1);`,
		`INSERT INTO "Status" ("Label", "Icon", "Notes", "IsLocked") VALUES ("Invading", "HelpCircle", NULL, 1);`,
		`INSERT INTO "DeviceTypes" ("Label", "Icon", "Notes", "IsLocked") VALUES ("Unspecified", "HelpCircle", NULL, 1);`,
		`INSERT INTO "InterfaceTypes" ("Label",  "Icon", "Notes", "IsLocked") VALUES ("WiFi", "HelpCircle", NULL, 1), ("Ethernet Cable", "HelpCircle", NULL, 1), ("Fibre", "HelpCircle", NULL, 1), ("Internal", "HelpCircle", "For MACVLAN containers, Virtual interfaces, etc", 1);`,
		`INSERT INTO "VLANs" ("Label", "Notes", "IPv4Mask", "IPv6Mask", "IsLocked") VALUES ("Default", NULL, "0.0.0.0", "", 1);`,
		`INSERT INTO "Locations" ("Label", "IsCloud", "Notes", "IsLocked") VALUES ("Limbo", 0, NULL, 1);`,
		`INSERT INTO "OperatingSystems" ("Vendor", "Family", "Version", "Name", "IsOpenSource", "IsServer", "Notes", "IsLocked") VALUES ("?", "?", "?", "?", 0, 0, NULL, 1), ("Apple", "MacOS", "10.13", "High Sierra", 0, 0, NULL, 0), ("Apple", "MacOS", "10.14", "Mojave", 0, 0, NULL, 0), ("Apple", "MacOS", "10.15", "Catalina", 0, 0, NULL, 0), ("Apple", "MacOS", "11", "Big Sur", 0, 0, NULL, 0), ("Apple", "MacOS", "12", "Monterey", 0, 0, NULL, 0), ("Microsoft", "Windows", "10", "Win10", 0, 0, NULL, 0), ("Microsoft", "Windows", "11", "Win11", 0, 0, NULL, 0), ("Canonical", "Ubuntu", "18.04", "Bionic Beaver", 1, 0, NULL, 0), ("Canonical", "Ubuntu", "20.04", "Focal Fossa", 1, 0, NULL, 0), ("Canonical", "Ubuntu", "21.04", "Hirsute Hippo", 1, 0, NULL, 0), ("Canonical", "Ubuntu", "21.10", "Impish Indri", 1, 0, NULL, 0), ("Canonical", "Ubuntu", "22.04", "Jammy Jellyfish", 1, 0, NULL, 0), ("Google", "Android", "6", "Marshmallow", 1, 0, NULL, 0), ("Google", "Android", "7", "Nougat", 1, 0, NULL, 0), ("Google", "Android", "8", "Oreo", 1, 0, NULL, 0), ("Google", "Android", "9", "Pie", 1, 0, NULL, 0), ("Google", "Android", "10", "Q", 1, 0, NULL, 0), ("Google", "Android", "11", "R", 1, 0, NULL, 0), ("Google", "Android", "12", "S", 1, 0, NULL, 0), ("Google", "Android", "13", "T", 1, 0, NULL, 0);`,
		`INSERT INTO "Architectures" ("Label", "BitSpace", "Notes", "IsLocked") VALUES ("Unknown", 0, NULL, 1), ("x86", 32, NULL, 0), ("x64", 64, NULL, 0), ("ARM", 32, NULL, 0), ("ARM64", 64, NULL, 0), ("RISC-V 32I", 32, NULL, 0), ("RISC-V 32E", 32, NULL, 0), ("RISC-V 64I", 64, NULL, 0), ("RISC-V 128I", 128, NULL, 0);`,

		fmt.Sprintf(`INSERT INTO "Users" ("Username", "Password", "Label", "CanAuthenticate", "AccessLevel", "IsInternal", "Notes", "IsLocked") VALUES (NULL, NULL, "Invader", 0, 0, 0, NULL, 1), ("Admin", "%s", "Admin", 1, 1, 0, NULL, 1);`, newPassword),

		`INSERT INTO "Meta" ("Name", "Value") VALUES ("DBVersion", "1");`,
	}

	for _, query = range create {
		_, err = Conn.Exec(query)
		if err != nil {
			return
		}
	}

	for _, query = range insert {
		_, err = Conn.Exec(query)
		if err != nil {
			return
		}
	}
	return
}
