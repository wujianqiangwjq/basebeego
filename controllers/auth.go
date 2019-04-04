package controllers

import (
	"apilogin/util"
	"strings"

	"github.com/astaxie/beego/context"
)

func Auth(ctx *context.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	token_total := strings.Fields(auth)

	if len(token_total) == 0 {
		if ctx.Request.URL.Path != "/v1/login" {
			ctx.Output.JSON(map[string]string{"errmsg": "invalid user/password"}, true, true)
		}

	} else if len(token_total) == 2 {
		if ok, username := util.ParseToken(token_total[1]); !ok && token_total[0] == "Jwt" {
			ctx.Output.JSON(map[string]string{"errmsg": "invalid user/password"}, true, true)
		} else {
			ctx.Input.SetData("user", username)
		}

	} else {
		ctx.Output.JSON(map[string]string{"errmsg": "invalid user/password"}, true, true)
	}

}
