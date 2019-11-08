package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {

	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}
	password := os.Getenv("MYSQL_PASSWD")
	username := os.Getenv("MYSQL_USERNAME")
	if username == "" {
		username = "root"
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}
	dbname := os.Getenv("MYSQL_DBNAME")
	if dbname == "" {
		dbname = "mysql"
	}
	data := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, dbname)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(localhost:3306)/wujq?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, false)

}
