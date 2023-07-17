package postgresql_test

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/gotabit/gotabit/app/params"
	"github.com/stretchr/testify/suite"

	"github.com/gotabit/juno/v5/database"
	databaseconfig "github.com/gotabit/juno/v5/database/config"
	postgres "github.com/gotabit/juno/v5/database/postgresql"
	"github.com/gotabit/juno/v5/logging"
)

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

type DbTestSuite struct {
	suite.Suite

	database *postgres.Database
}

func (suite *DbTestSuite) SetupTest() {
	// Create the codec
	codec := params.MakeEncodingConfig()

	// Build the database
	dbCfg := databaseconfig.NewDatabaseConfig(
		"postgres://bdjuno:password@localhost:6433/bdjuno?sslmode=disable&search_path=public",
		"false",
		"",
		"",
		"",
		-1,
		-1,
		100000,
		100,
	)
	db, err := postgres.Builder(database.NewContext(dbCfg, &codec, logging.DefaultLogger()))
	suite.Require().NoError(err)

	bigDipperDb, ok := (db).(*postgres.Database)
	suite.Require().True(ok)

	// Delete the public schema
	_, err = bigDipperDb.SQL.Exec(`DROP SCHEMA public CASCADE;`)
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = bigDipperDb.SQL.Exec(`CREATE SCHEMA public;`)
	suite.Require().NoError(err)

	dirPath := path.Join(".")
	dir, err := ioutil.ReadDir(dirPath)
	suite.Require().NoError(err)

	for _, fileInfo := range dir {
		if !strings.Contains(fileInfo.Name(), ".sql") {
			continue
		}

		file, err := ioutil.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		suite.Require().NoError(err)

		commentsRegExp := regexp.MustCompile(`/\*.*\*/`)
		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err := bigDipperDb.SQL.Exec(commentsRegExp.ReplaceAllString(request, ""))
			suite.Require().NoError(err)
		}
	}

	suite.database = bigDipperDb
}
