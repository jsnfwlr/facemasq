package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"facemasq/lib/files"
	"facemasq/lib/logging"
	"facemasq/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

const DBTargetVer = 1

var (
	adminPassword string
	DBEngine      string
	DBConnString  string
	Conn          *bun.DB
	Context       context.Context
	Verbose       bool
)

func init() {
	adminPassword = os.Getenv("ADMINPASSWORD")
	dbConnString := os.Getenv("DBCONNSTR")
	if dbConnString == "" {
		dbConnString = "sqlite://network.sqlite"
	}
	err := initialise(dbConnString)
	if err != nil {
		logging.Panic(err)
	}
}

func initialise(dbConnString string) (err error) {
	var dataPath string
	DBEngine = ""
	if !strings.Contains(dbConnString, "://") {
		err = fmt.Errorf("`%s` is in the wrong format", dbConnString)
		return
	}
	dbParams := strings.Split(dbConnString, "://")
	DBEngine = strings.ToLower(dbParams[0])

	switch DBEngine {
	case "sqlite":
		dataPath, err = files.GetDir("data")
		if err != nil {
			return
		}
		DBConnString = fmt.Sprintf("file:%[2]s%[1]c%[3]s", os.PathSeparator, dataPath, dbParams[1])

	case "mariadb", "mysql":
		DBConnString = fmt.Sprintf("%s?parseTime=true", dbParams[1])
	case "pg", "pgsql", "postgres", "postgresql":
		DBEngine = "postgres"
		DBConnString = fmt.Sprintf("postgres://%s?sslmode=disable", dbParams[1])
	default:
		err = fmt.Errorf("`%s` is not supported", DBEngine)
	}

	return
}

func Connect() (err error) {
	var conn *sql.DB
	var doPrepare bool

	Context = context.Background()

	switch DBEngine {
	case "sqlite":
		conn, err = sql.Open(sqliteshim.ShimName, DBConnString)
		if err != nil {
			logging.Panic(err)
		}
		conn.SetMaxOpenConns(1)
		Conn = bun.NewDB(conn, sqlitedialect.New())
	case "mariadb", "mysql":
		conn, err = sql.Open("mysql", DBConnString)
		if err != nil {
			logging.Panic(err)
		}
		Conn = bun.NewDB(conn, mysqldialect.New())
	case "postgres":
		DBEngine = "postgres"
		config, err := pgx.ParseConfig(DBConnString)
		if err != nil {
			logging.Panic(err)
		}
		config.PreferSimpleProtocol = true

		conn := stdlib.OpenDB(*config)
		Conn = bun.NewDB(conn, pgdialect.New())
	}

	Conn.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("BUNDEBUG"),
	))
	doPrepare, err = checkPrepare()
	if doPrepare {
		if err := prepare(Context, Conn); err != nil {
			logging.Panic(err)
		}
	}
	return
}

func checkPrepare() (doPrepare bool, err error) {
	var sql string
	if strings.HasSuffix(os.Args[0], ".test") {
		doPrepare = true
		return
	}

	doPrepare = false
	switch DBEngine {
	case "sqlite":
		sql = `SELECT name FROM sqlite_master WHERE type='table' AND name='meta';`
	case "mysql":
		sql = `SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'meta';`
	case "postgres":
		sql = `SELECT table_name FROM information_schema.tables WHERE table_name = 'meta';`
	}
	var name string
	err = Conn.NewRaw(sql).Scan(Context, &name)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return
		}
		doPrepare = true
		err = nil
		return
	}

	return
}

func prepare(ctx context.Context, db *bun.DB) (err error) {
	logging.Info("Preparing db")
	var drops, creates []interface{}
	// Params
	drops = append(
		drops,

		// Settings
		(*models.Meta)(nil),

		// Scans
		(*models.Port)(nil),
		(*models.Scan)(nil),
		(*models.History)(nil),

		// Devices
		(*models.Hostname)(nil),
		(*models.Address)(nil),
		(*models.Interface)(nil),
		(*models.Device)(nil),

		// Users & Maintainers
		(*models.User)(nil),

		// Params
		(*models.VLAN)(nil),
		(*models.Status)(nil),
		(*models.OperatingSystem)(nil),
		(*models.Location)(nil),
		(*models.InterfaceType)(nil),
		(*models.DeviceType)(nil),
		(*models.Category)(nil),
		(*models.Architecture)(nil),
	)
	creates = append(
		creates,

		// Params
		(*models.Architecture)(nil),
		(*models.Category)(nil),
		(*models.DeviceType)(nil),
		(*models.InterfaceType)(nil),
		(*models.Location)(nil),
		(*models.OperatingSystem)(nil),
		(*models.Status)(nil),
		(*models.VLAN)(nil),

		// Users & Maintainers
		(*models.User)(nil),

		// Devices
		(*models.Device)(nil),
		(*models.Interface)(nil),
		(*models.Address)(nil),
		(*models.Hostname)(nil),

		// Scans
		(*models.History)(nil),
		(*models.Scan)(nil),
		(*models.Port)(nil),

		// Settings
		(*models.Meta)(nil),
	)
	for _, drop := range drops {
		if _, err = db.NewDropTable().Model(drop).IfExists().Cascade().Exec(ctx); err != nil {
			return
		}
	}

	if err = db.ResetModel(ctx, creates...); err != nil {
		return
	}
	userSeed := models.GetUserSeed(adminPassword)
	if _, err = db.NewInsert().Model(&userSeed).Exec(ctx); err != nil {
		return
	}
	architectureSeed := models.GetArchitectureSeed()
	if _, err = db.NewInsert().Model(&architectureSeed).Exec(ctx); err != nil {
		return
	}
	categorySeed := models.GetCategorySeed()
	if _, err = db.NewInsert().Model(&categorySeed).Exec(ctx); err != nil {
		return
	}
	deviceTypeSeed := models.GetDeviceTypeSeed()
	if _, err = db.NewInsert().Model(&deviceTypeSeed).Exec(ctx); err != nil {
		return
	}
	interfaceTypeSeed := models.GetInterfaceTypeSeed()
	if _, err = db.NewInsert().Model(&interfaceTypeSeed).Exec(ctx); err != nil {
		return
	}
	locationSeed := models.GetLocationSeed()
	if _, err = db.NewInsert().Model(&locationSeed).Exec(ctx); err != nil {
		return
	}
	operatingSystemSeed := models.GetOperatingSystemSeed()
	if _, err = db.NewInsert().Model(&operatingSystemSeed).Exec(ctx); err != nil {
		return
	}
	statusSeed := models.GetStatusSeed()
	if _, err = db.NewInsert().Model(&statusSeed).Exec(ctx); err != nil {
		return
	}
	vLANSeed := models.GetVLANSeed()
	if _, err = db.NewInsert().Model(&vLANSeed).Exec(ctx); err != nil {
		return
	}

	return
}

// func ConnectToTest() (err error) {
// 	return
// }
