package controllers

import (
	"basebeego/models"
	"basebeego/util"
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (user *UserController) Post() {
	var params map[string]interface{}
	json.Unmarshal(user.Ctx.Input.RequestBody, &params)
	name, ok := params["name"]
	if !ok {
		user.Data["json"] = map[string]string{"status": "failed", "ermsg": "parameter is wrong"}
		user.ServeJSON()
		return
	}
	passwd, ok := params["passwd"]
	if !ok {
		user.Data["json"] = map[string]string{"status": "failed", "ermsg": "parameter is wrong"}
		user.ServeJSON()
		return
	}
	role := 1
	role_tem, ok := params["role"]
	if ok {
		role = int(role_tem.(float64))
	}

	usermode := models.User{Name: name.(string)}
	flag_exist := usermode.CheckUserExistByName()
	if flag_exist {
		user.Data["json"] = map[string]string{"status": "failed", "ermsg": "user exists"}
		user.ServeJSON()

	} else {
		usermode.Role = role
		usermode.Paswd = util.Md5Password(passwd.(string))
		usermode.CreateTime = time.Now().Unix()
		usermode.Status = 0
		flag := models.AddUser(&usermode)
		if flag {
			user.Data["json"] = map[string]string{"status": "success"}

		} else {
			user.Data["json"] = map[string]string{"status": "failed"}
		}
		user.ServeJSON()
	}

}

func (user *UserController) Get() {
	flag, _ := models.GetUsers()
	data, _ := json.Marshal(flag)

	user.Data["json"] = map[string]string{"status": "success", "data": string(data)}
	user.ServeJSON()

}

func (user *UserController) Put() {
	param, _ := strconv.Atoi(user.Ctx.Input.Param(":id"))
	user_mode := models.User{Id: param}
	user_mode.DeleteUser()
	user.Data["json"] = map[string]string{"status": "success"}
	user.ServeJSON()
}
