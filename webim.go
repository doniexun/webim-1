package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/webim/api"
	"github.com/adolphlwq/webim/service"
)

var (
	dbname      string
	user        string
	pass        string
	dbaddress   string
	dbport      string
	serviceport string
	serviceaddr string
)

func main() {
	flag.StringVar(&dbname, "dbname", "tinyurl", "database name to connect")
	flag.StringVar(&user, "user", "test", "user of database")
	flag.StringVar(&pass, "pass", "test", "pass of database")
	flag.StringVar(&dbaddress, "dbaddress", "localhost", "address of database")
	flag.StringVar(&dbport, "dbport", "3306", "port of database")
	flag.StringVar(&serviceport, "port", "8877", "port tinyurl bind on")
	flag.StringVar(&serviceaddr, "serviceaddr", "0.0.0.0", "address webim bind on")
	flag.Parse()

	logrus.Info("Start init DB")
	dbs := service.NewDB(dbname, user, pass, dbaddress, dbport)
	api.WebIMAPI(serviceaddr+":"+serviceport, dbs)
}
