package controllers

import (
	"github.com/astaxie/beego"
)

type LiveController struct {
	beego.Controller
}

func (live *LiveController) Get() {
	live.Data["json"] = map[string]string{"status": "ok", "data": "test"}
	live.ServeJSON()
}
