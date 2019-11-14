package controllers

import (
	"strings"

	"../util"

	"github.com/astaxie/beego/context"
)

func skip_auth(path string) bool {
	skip_path := []string{
		"/v1/login",
		"/v1/user",
		"/v1",
	}
	for _, value := range skip_path {
		if (path == value) || (path == (value + "/")) {
			return true
		}
	}
	return false

}
func Auth(ctx *context.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	token_total := strings.Fields(auth)

	if len(token_total) == 0 {
		if !skip_auth(ctx.Request.URL.Path) {
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
