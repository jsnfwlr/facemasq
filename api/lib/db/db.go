package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"facemasq/lib/files"
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
)

func init() {
	var err error
	var dataPath string

	adminPassword = os.Getenv("ADMINPASSWORD")
	DBConnString = os.Getenv("DBCONNSTR")
	if DBConnString == "" {
		DBConnString = "sqlite://network.sqlite"
	}

	dbParams := strings.Split(DBConnString, "://")
	DBEngine = strings.ToLower(dbParams[0])

	switch DBEngine {
	case "sqlite":
		dataPath, err = files.GetDir("data")
		if err != nil {
			panic(err)
		}
		DBConnString = fmt.Sprintf("file:%[2]s%[1]c%[3]s", os.PathSeparator, dataPath, dbParams[1])

	case "mariadb", "mysql":
		DBConnString = fmt.Sprintf("%s?parseTime=true", dbParams[1])
	case "pg", "pgsql", "postgres", "postgresql":
		DBEngine = "postgres"
		DBConnString = fmt.Sprintf("postgres://%s?sslmode=disable", dbParams[1])
	}
}

func Connect() (err error) {
	var conn *sql.DB
	var doPrepare bool

	Context = context.Background()

	switch DBEngine {
	case "sqlite":
		conn, err = sql.Open(sqliteshim.ShimName, DBConnString)
		if err != nil {
			panic(err)
		}
		conn.SetMaxOpenConns(1)
		Conn = bun.NewDB(conn, sqlitedialect.New())
	case "mariadb", "mysql":
		conn, err = sql.Open("mysql", DBConnString)
		if err != nil {
			panic(err)
		}
		Conn = bun.NewDB(conn, mysqldialect.New())
	case "postgres":
		config, err := pgx.ParseConfig(DBConnString)
		if err != nil {
			panic(err)
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
			panic(err)
		}
	}
	return
}

func checkPrepare() (doPrepare bool, err error) {
	var sql string

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
	doPrepare = false
	return
}

func prepare(ctx context.Context, db *bun.DB) error {
	log.Println("Preparing db")
	if err := db.ResetModel(ctx,
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
		(*models.History)(nil),

		// Scans
		(*models.Scan)(nil),
		(*models.Port)(nil),
		(*models.Meta)(nil),
	); err != nil {
		return err
	}
	userSeed := models.GetUserSeed(adminPassword)
	if _, err := db.NewInsert().Model(&userSeed).Exec(ctx); err != nil {
		return err
	}
	architectureSeed := models.GetArchitectureSeed()
	if _, err := db.NewInsert().Model(&architectureSeed).Exec(ctx); err != nil {
		return err
	}
	categorySeed := models.GetCategorySeed()
	if _, err := db.NewInsert().Model(&categorySeed).Exec(ctx); err != nil {
		return err
	}
	deviceTypeSeed := models.GetDeviceTypeSeed()
	if _, err := db.NewInsert().Model(&deviceTypeSeed).Exec(ctx); err != nil {
		return err
	}
	interfaceTypeSeed := models.GetInterfaceTypeSeed()
	if _, err := db.NewInsert().Model(&interfaceTypeSeed).Exec(ctx); err != nil {
		return err
	}
	locationSeed := models.GetLocationSeed()
	if _, err := db.NewInsert().Model(&locationSeed).Exec(ctx); err != nil {
		return err
	}
	operatingSystemSeed := models.GetOperatingSystemSeed()
	if _, err := db.NewInsert().Model(&operatingSystemSeed).Exec(ctx); err != nil {
		return err
	}
	statusSeed := models.GetStatusSeed()
	if _, err := db.NewInsert().Model(&statusSeed).Exec(ctx); err != nil {
		return err
	}
	vLANSeed := models.GetVLANSeed()
	if _, err := db.NewInsert().Model(&vLANSeed).Exec(ctx); err != nil {
		return err
	}

	return nil
}
