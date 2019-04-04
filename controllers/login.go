package controllers

import (
	"apilogin/models"
	"apilogin/util"
	"encoding/json"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (login *LoginController) Post() {
	var user map[string]string
	json.Unmarshal(login.Ctx.Input.RequestBody, &user)
	username, ok := user["name"]
	if !ok {
		login.Data["json"] = map[string]string{"status": "failed"}
		login.ServeJSON()
	}

	password, ok := user["passwd"]
	if !ok {
		login.Data["json"] = map[string]string{"status": "failed"}
		login.ServeJSON()
	}
	var usermode models.User
	usermode.Name = username
	dbuser, err := models.GetUser(&usermode)
	if err == nil {
		if dbuser.Paswd == util.Md5Password(password) {
			if token, err := util.GetToken(username); err == nil {
				login.Data["json"] = map[string]string{"status": "success", "token": token}
				login.ServeJSON()
			}
		}
	}
	login.Data["json"] = map[string]string{"status": "failed"}
	login.ServeJSON()

}
