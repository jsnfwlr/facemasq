package db

import "testing"

type DBTestSet struct {
	ConnString     string
	ExpectedEngine string
	ExpectedError  string
}

func TestDBInit(t *testing.T) {
	dbTestSet := []DBTestSet{
		{
			ConnString:     "sqlite://network.sql",
			ExpectedEngine: "sqlite",
			ExpectedError:  "",
		},
		{
			ConnString:     "MariaDB://username:password@(host:port)/database",
			ExpectedEngine: "mariadb",
			ExpectedError:  "",
		},
		{
			ConnString:     "mysql://username:password@(host:port)/database",
			ExpectedEngine: "mysql",
			ExpectedError:  "",
		},
		{
			ConnString:     "pg://username:password@host:port/database?sslmode=disable",
			ExpectedEngine: "postgres",
			ExpectedError:  "",
		},
		{
			ConnString:     "postgres://username:password@host:port/database?sslmode=disable",
			ExpectedEngine: "postgres",
			ExpectedError:  "",
		},
		{
			ConnString:     "postgresql://username:password@host:port/database?sslmode=disable",
			ExpectedEngine: "postgres",
			ExpectedError:  "",
		},
		{
			ConnString:     "pgsql://username:password@host:port/database?sslmode=disable",
			ExpectedEngine: "postgres",
			ExpectedError:  "",
		},
		{
			ConnString:     "mssql://username:password@host:port/database",
			ExpectedEngine: "mssql",
			ExpectedError:  "`mssql` is not supported",
		},
		{
			ConnString:     "",
			ExpectedEngine: "",
			ExpectedError:  "`` is in the wrong format",
		},
	}
	for i := range dbTestSet {
		err := initialise(dbTestSet[i].ConnString)
		if err != nil && err.Error() != dbTestSet[i].ExpectedError {
			t.Error(err)
			continue
		}
		if DBEngine != dbTestSet[i].ExpectedEngine {
			t.Errorf("DBEngine was expected to be `%s` - got `%s`", dbTestSet[i].ExpectedEngine, DBEngine)
		}
	}
}
