package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func Init() {
	host := beego.AppConfig.String("db.host")
	password := beego.AppConfig.String("db.passwd")
	username := beego.AppConfig.String("db.username")
	port := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.dbname")
	data := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		username,
		password,
		dbname,
		host,
		port)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", data)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, false)

}
