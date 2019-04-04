package controllers

import (
	"apilogin/models"
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
	name := params["name"].(string)
	passwd := params["passwd"].(string)
	role := int(params["role"].(float64))
	usermode := models.User{Name: name, Paswd: passwd, Role: role}
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

func (user *UserController) Get() {
	flag, _ := models.GetUsers()
	data, _ := json.Marshal(flag)

	user.Data["json"] = map[string]string{"status": "success", "data": string(data)}
	user.ServeJSON()

}

func (user *UserController) Put() {
	param, _ := strconv.Atoi(user.Ctx.Input.Param(":id"))
	models.DeleteUser(&models.User{Id: param})
	user.Data["json"] = map[string]string{"status": "success"}
	user.ServeJSON()
}
