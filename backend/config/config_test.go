package config_test

import (
	"github.com/MisterZurg/TBank-backend-academy-URL-Shortener/backend/config"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

var ConfigEnvs = []string{
	"APP_PORT",
	"REDIS_HOST", "REDIS_PORT",
	"CLICKHOUSE_USER", "CLICKHOUSE_PASSWORD", "CLICKHOUSE_DB", "CLICKHOUSE_HOST", "CLICKHOUSE_PORT",
}

// os.Setenv("FOO", "1")
// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context.
type ConfigSuite struct {
	suite.Suite

	cfg *config.Config
}

// Executes before each test case.
func (suite *ConfigSuite) SetupTest() {}

// Executes after each test case.
func (suite *ConfigSuite) TearDownTest() {
	for _, e := range ConfigEnvs {
		os.Unsetenv(e)
	}

	suite.cfg = nil
}

// TestSuiteCerts runs all suite tests.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) TestNewConfig() {
	tests := map[string]struct {
		valuesToSet map[string]string
		expCfg      config.Config
	}{
		"Default": {
			valuesToSet: map[string]string{},
			expCfg: config.Config{
				APPPort:    "1323",
				REDISHost:  "localhost",
				REDISPort:  "6379",
				CHUser:     "oleg",
				CHPassword: "tinkoff",
				CHDBName:   "tbank_academy",
				CHHost:     "localhost",
				CHPort:     "19000",
			},
		},
		"With Custom Values": {
			valuesToSet: map[string]string{
				"REDIS_HOST":      "barak.obama",
				"REDIS_PORT":      "prigozhin",
				"CLICKHOUSE_DB":   "mr_ildar",
				"CLICKHOUSE_HOST": "sir",
				"CLICKHOUSE_PORT": "emil",
			},
			expCfg: config.Config{
				APPPort:    "1323",
				REDISHost:  "barak.obama",
				REDISPort:  "prigozhin",
				CHUser:     "oleg",
				CHPassword: "tinkoff",
				CHDBName:   "mr_ildar",
				CHHost:     "sir",
				CHPort:     "emil",
			},
		},
	}

	for name, test := range tests {
		suite.T().Run(name, func(t *testing.T) {
			for envKey, envVal := range test.valuesToSet {
				os.Setenv(envKey, envVal)
			}

			gotCfg, err := config.New()
			suite.Require().NoError(err)

			suite.Require().Equal(test.expCfg.APPPort, gotCfg.APPPort, "App ports, doesn't match")
			suite.Require().Equal(test.expCfg.GetAppAddress(), gotCfg.GetAppAddress(), "App GetAppAddress(), doesn't match")

			suite.Require().Equal(test.expCfg.REDISHost, gotCfg.REDISHost, "REDISHost, doesn't match")
			suite.Require().Equal(test.expCfg.REDISPort, gotCfg.REDISPort, "REDISPort, doesn't match")
			suite.Require().Equal(test.expCfg.GetRedisDSN(), gotCfg.GetRedisDSN(), "GetRedisDSN, doesn't match")

			suite.Require().Equal(test.expCfg.CHUser, gotCfg.CHUser, "CHUser, doesn't match")
			suite.Require().Equal(test.expCfg.CHHost, gotCfg.CHHost, "CHHost, doesn't match")
			suite.Require().Equal(test.expCfg.CHPort, gotCfg.CHPort, "CHPort, doesn't match")
			suite.Require().Equal(test.expCfg.CHPassword, gotCfg.CHPassword, "CHPassword, doesn't match")
			suite.Require().Equal(test.expCfg.CHDBName, gotCfg.CHDBName, "CHDBName, doesn't match")
			suite.Require().Equal(test.expCfg.GetClickHouseDSN(), gotCfg.GetClickHouseDSN(), "GetClickHouseDSN(), doesn't match")

		})
	}
}
