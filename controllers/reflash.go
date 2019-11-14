package controllers

import (
	"../util"
	"github.com/astaxie/beego"
)

type TokenController struct {
	beego.Controller
}

func (r *TokenController) Post() {
	data := r.Ctx.Input.Data()
	token, err := util.GetToken(data["user"].(string))
	if err == nil {
		r.Data["json"] = map[string]string{"token": token, "status": "success"}
		r.ServeJSON()
	} else {
		r.Data["json"] = map[string]string{"status": "failed"}
		r.ServeJSON()
	}
}
