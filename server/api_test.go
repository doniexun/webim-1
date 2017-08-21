package server

import (
	"testing"
	"time"

	"github.com/adolphlwq/webim/mysql"
)

const (
	TestPort = "9066"
	TestAddr = "http://0.0.0.0:9066"
)

var (
	mysqlClient *mysql.Client
	appService  *ServiceProvider
	configPath  string
)

func init() {
	configPath = "../test.properties"
	mysqlClient = mysql.NewMySQLClient(configPath)
	appService = &ServiceProvider{
		MysqlClient: mysqlClient,
	}
}

func clearDatabase() {
	appService.MysqlClient.DropDatabase()
}

func StartTestServer(t *testing.T) {
	go func() {
		testServer := builfEngine(appService)
		testServer.Run(TestAddr)
	}()

	t.Logf("wait 2s for test server to start...\n")
	time.Sleep(time.Second * 2)
}
