package mysql

import (
	"database/sql"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Client hold mysql connection and wrape CRUD methods
type Client struct {
	DB       *gorm.DB
	Database string
}

// NewMySQLClient return new MySQLClient instance
func NewMySQLClient(configPath string) *Client {
	setting := NewMySQLConfig(configPath)
	mc := &Client{}

	CheckDB(setting)
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.User, setting.Password, setting.Host, setting.Port, setting.Database)
	db, err := gorm.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("open connection to mysql use gorm error: %s", err)
	}

	mc.Database = setting.Database
	mc.DB = db
	mc.DB.AutoMigrate(&entity.User{}, &entity.Contact{}, &entity.Message{})
	logrus.Infof("create table users, contacts and messages done.\n")

	return mc
}

// CheckDB check if database existed in db
// create database if not
func CheckDB(setting *MySQLConfig) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/", setting.User, setting.Password,
		setting.Host, setting.Port)
	db, err := sql.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("connection to db error: %s", err)
	}
	defer db.Close()

	sql := "CREATE DATABASE IF NOT EXISTS " + setting.Database + ";"
	_, err = db.Exec(sql)
	if err != nil {
		logrus.Fatalf("create db %s error: %v", setting.Database, err)
	}
}

// DropDatabase drop self hold database
func (c *Client) DropDatabase() {
	sql := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", c.Database)
	db := c.DB.DB()

	_, err := db.Exec(sql)
	if err != nil {
		logrus.Fatalf("drop database %s error: %v", c.Database, err)
	}
}

// ChekcUserExistedByUsername check table users to see if username existed
// true existed and false not
func (c *Client) ChekcUserExistedByUsername(username string) bool {
	var tmpUser entity.User
	c.DB.Where("username = ?", username).First(&tmpUser)
	return tmpUser.Username != ""
}

// truncateTable truncate table
func (c *Client) truncateTable(table string) {
	sqlDB := c.DB.DB()
	sql := "TRUNCATE TABLE " + table

	_, err := sqlDB.Exec(sql)
	if err != nil {
		logrus.Fatalf("truncate table %s error %v.\n", table, err)
	}
}
