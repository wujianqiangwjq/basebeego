package controllers

import (
	"encoding/json"

	"io/ioutil"

	"../models"
	"../util"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (login *LoginController) Post() {
	var user map[string]string
	logdata, _ := ioutil.ReadAll(login.Ctx.Request.Body)
	jsonerr := json.Unmarshal(logdata, &user)
	if jsonerr != nil {
		login.Data["json"] = map[string]string{"status": "failed", "ermsg": jsonerr.Error()}
		login.ServeJSON()
		return
	}
	username, ok := user["name"]
	if !ok {
		login.Data["json"] = map[string]string{"status": "failed", "ermsg": "can not get name parameter"}
		login.ServeJSON()
		return
	}

	password, ok := user["passwd"]
	if !ok {
		login.Data["json"] = map[string]string{"status": "failed", "ermsg": "can not get passwd parameter"}
		login.ServeJSON()
		return
	}
	var usermode models.User
	usermode.Name = username
	dbuser, err := usermode.GetUser("Name")
	if err == nil {
		if dbuser.Paswd == util.Md5Password(password) {
			if token, err := util.GetToken(username); err == nil {
				login.Data["json"] = map[string]string{"status": "success", "token": token}
				login.ServeJSON()
				return
			}
		}
	}
	login.Data["json"] = map[string]string{"status": "failed", "ermsg": "username/password is wrong"}
	login.ServeJSON()

}
