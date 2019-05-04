package main

import (
	"basebeego/models"
	_ "basebeego/routers"
	"basebeego/util"

	"github.com/astaxie/beego"
)

func main() {
	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	// beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	beego.AppConfig.Set("token_key", util.GetGernate())
	models.Init()
	beego.Run()
}
