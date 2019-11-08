package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	// host := beego.AppConfig.String("db.host")
	// password := beego.AppConfig.String("db.passwd")
	// username := beego.AppConfig.String("db.username")
	// port := beego.AppConfig.String("db.port")
	// dbname := beego.AppConfig.String("db.dbname")
	// data := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
	// 	username,
	// 	password,
	// 	dbname,
	// 	host,
	// 	port)
	// data := ""
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(localhost:3306)/wujq?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, false)

}
