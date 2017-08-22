package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	"github.com/adolphlwq/webim/mysql"
	"github.com/adolphlwq/webim/util"
	"github.com/gin-gonic/gin"
)

const (
	TestPort     = "9066"
	TestAddr     = "http://0.0.0.0:9066"
	TestUsername = "test"
	TestPassword = "123456"
)

var (
	mysqlClient     *mysql.Client
	appService      *ServiceProvider
	configPath      string
	encryptPassword string
)

func init() {
	encryptPassword = util.EncryptPasswordWithSalt(TestUsername, TestPassword)
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
	logrus.Println("delete test database ", appService.MysqlClient.Database)
}

func startTestServer(t *testing.T) {
	go func() {
		testServer := builfEngine(appService)
		testServer.Run(":" + TestPort)
	}()

	t.Logf("wait 2s for test server to start...\n")
	time.Sleep(time.Second * 2)
}

func postJSON(postURL string, data io.Reader) map[string]interface{} {
	resp, err := http.Post(postURL, "application/json", data)
	if err != nil {
		logrus.Fatalf("post data error: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("read body of response error: %v", err)
	}

	var mp map[string]interface{}
	err = json.Unmarshal(body, &mp)
	if err != nil {
		logrus.Fatalf("parse body of response error: %v", err)
	}

	return mp
}

func insertTestUser(user entity.User) {
	appService.MysqlClient.DB.Create(&user)
}
