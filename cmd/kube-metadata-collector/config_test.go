package main

import (
	"os"
	"testing"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
    suite.Suite
}

func (suite *ConfigTestSuite) SetupTest() {
	os.Unsetenv("PILLY_DB_URI")
	os.Unsetenv("PILLY_INTERVAL")
}

func (suite *ConfigTestSuite) TestGetConfigValid() {
	os.Setenv("PILLY_DB_URI", "sqlite://database.sql")
	testConfig := GetConfig()
	suite.Equal(testConfig.DbURI, "sqlite://database.sql")
	suite.Equal(testConfig.Interval, 60)
}

func (suite *ConfigTestSuite) TestGetConfigValidSpecificInterval() {
	os.Setenv("PILLY_DB_URI", "sqlite://database.sql")
	os.Setenv("PILLY_INTERVAL", "600")
	testConfig := GetConfig()
	suite.Equal(testConfig.DbURI, "sqlite://database.sql")
	suite.Equal(testConfig.Interval, 600)
}

/*func (suite *ConfigTestSuite) TestGetConfigInValidNoDbURI() {
	testConfig := GetConfig()
	suite.Equal(testConfig.Interval, 60)
}*/

func TestRunSuite(t *testing.T) {
    suite.Run(t, new(ConfigTestSuite))
}