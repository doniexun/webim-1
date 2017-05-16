package service

import (
	"database/sql"
	"sync"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

var onceSetupDB sync.Once

type DBService struct {
	DBName  string
	User    string
	Pass    string
	Address string
	Port    string
	DBPath  string
}

// NewDB return *DBService
func NewDB(dbname, user, pass, address, port string) *DBService {
	dbpath := user + ":" + pass + "@tcp(" + address + ":" + port + ")/"
	logrus.Info("dbpath is ", dbpath)

	dbs := &DBService{
		DBName:  dbname,
		User:    user,
		Pass:    pass,
		Address: address,
		Port:    port,
		DBPath:  dbpath,
	}

	dbs.Setup()

	return dbs
}

// CreateBareDB create *sql.DB without connecting to database name
func (dbs *DBService) CreateBareDB() *sql.DB {
	db, err := sql.Open("mysql", dbs.DBPath)
	if err != nil {
		logrus.Fatal("setup up db error:", err)
	}
	logrus.Info("create bare db.")

	return db
}

// CreateDB create *sql.DB with connecting to database name
func (dbs *DBService) CreateDB() *sql.DB {
	db, err := sql.Open("mysql", dbs.DBPath+dbs.DBName)
	if err != nil {
		logrus.Fatal("setup up db error:", err)
	}
	logrus.Info("create db with dbname ", dbs.DBName)

	return db
}

// Setup connect to mysql and create database DBName if not exists
func (dbs *DBService) Setup() {
	db := dbs.CreateBareDB()
	defer db.Close()

	dbSchema := `
		CREATE DATABASE IF NOT EXISTS ` + dbs.DBName + ` 
			DEFAULT CHARACTER SET utf8mb4
			DEFAULT COLLATE utf8mb4_unicode_ci;
	`
	useDB := "USE " + dbs.DBName + ";"
	//utf8mb4 := "SET NAMES 'utf8mb4'; SET CHARACTER SET utf8mb4;"
	utf8mb4 := "SET NAMES 'utf8mb4';"

	userTable := `
		CREATE TABLE IF NOT EXISTS ` + dbs.DBName + `.user (
			id INT(64) NOT NULL AUTO_INCREMENT,
			username VARCHAR(20) NOT NULL,
			password VARBINARY(32) NOT NULL,
			created_time INT(64) NOT NULL,
			PRIMARY KEY (id)
		)
		CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
	`

	friendTable := `
		CREATE TABLE IF NOT EXISTS ` + dbs.DBName + `.friend_relationship (
			id INT(64) NOT NULL AUTO_INCREMENT,
			fmin VARCHAR(20) NOT NULL,
			fmax VARCHAR(20) NOT NULL,
			added_time INT(64) NOT NULL,
			state VARCHAR(10) NOT NULL,
			PRIMARY KEY (id)
		)
		CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
	`

	msgTable := `
		CREATE TABLE IF NOT EXISTS ` + dbs.DBName + `.message (
			id INT(64) NOT NULL AUTO_INCREMENT,
			sender VARCHAR(20) NOT NULL,
			receiver VARCHAR(20) NOT NULL,
			msg TEXT NOT NULL,
			send_time INT(64) NOT NULL,
			PRIMARY KEY (id)
		)
		CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
	`

	onceSetupDB.Do(func() {
		logrus.Info("start create db %s if not exists.", dbs.DBName)
		if _, err := db.Exec(dbSchema); err != nil {
			logrus.Fatal("setup database ", dbs.DBName, " err:", err)
		}

		logrus.Info("start use db ", dbs.DBName)
		if _, err := db.Exec(useDB); err != nil {
			logrus.Fatal("use db ", dbs.DBName, " error:", err)
		}

		logrus.Info("start create table user if not exists.")
		if _, err := db.Exec(userTable); err != nil {
			logrus.Fatal("setup table user error:", err)
		}

		logrus.Info("start create table friend_relationship if not exists.")
		if _, err := db.Exec(friendTable); err != nil {
			logrus.Fatal("setup table friend_relationship error:", err)
		}

		logrus.Info("start create table message if not exists.")
		if _, err := db.Exec(msgTable); err != nil {
			logrus.Fatal("setup table message error:", err)
		}

		logrus.Info("try use utf8mb4")
		if _, err := db.Exec(utf8mb4); err != nil {
			logrus.Fatal("use utf8mb4 error:", err)
		}
	})
}

// STMTFactory return *sql.Stmt and handle error
func (dbs *DBService) STMTFactory(predSQL string, db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare(predSQL)
	if err != nil {
		logrus.Fatalf("prepare sql %s error: %s", predSQL, err.Error())
	}
	return stmt
}
