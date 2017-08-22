package server

import (
	"testing"
	"time"

	"github.com/adolphlwq/webim/mysql"
	"github.com/gin-gonic/gin"
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

	// set gin mode
	gin.SetMode(gin.TestMode)
}

func clearDatabase() {
	appService.MysqlClient.DropDatabase()
}

func startTestServer(t *testing.T) {
	go func() {
		testServer := builfEngine(appService)
		testServer.Run(":" + TestPort)
	}()

	t.Logf("wait 2s for test server to start...\n")
	time.Sleep(time.Second * 2)
}
