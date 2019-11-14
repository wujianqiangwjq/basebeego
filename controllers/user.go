package controllers

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	"../models"
	"../util"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (user *UserController) Post() {
	var params map[string]string
	logdata, _ := ioutil.ReadAll(user.Ctx.Request.Body)
	jsonerr := json.Unmarshal(logdata, &params)
	if jsonerr != nil {
		user.Data["json"] = map[string]string{"status": "failed", "ermsg": jsonerr.Error()}
		user.ServeJSON()
		return
	}

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
		role, _ = strconv.Atoi(role_tem)
	}

	usermode := models.User{Name: name}
	flag_exist := usermode.CheckUserExistByName()
	if flag_exist {
		user.Data["json"] = map[string]string{"status": "failed", "ermsg": "user exists"}
		user.ServeJSON()

	} else {
		usermode.Role = role
		usermode.Paswd = util.Md5Password(passwd)
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
