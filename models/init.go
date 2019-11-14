package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {

	host := os.Getenv("MYSQLHOST")
	if host == "" {
		host = "localhost"
	}
	password := os.Getenv("MYSQLPASSWD")
	username := os.Getenv("MYSQLUSERNAME")
	if username == "" {
		username = "root"
	}
	port := os.Getenv("MYSQLPORT")
	if port == "" {
		port = "3306"
	}
	dbname := os.Getenv("MYSQLDBNAME")
	if dbname == "" {
		dbname = "mysql"
	}
	data := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, dbname)
	fmt.Println(data)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", data)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, false)

}
