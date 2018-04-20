package controllers

import (
	"github.com/astaxie/beego"
	"github.com/logoocc/tokenservice/token"
)

type TokenValidateController struct {
	beego.Controller
}

func (c *TokenValidateController) Get() {
	tokenString := c.GetString("token")
	var resultMap map[string]string
	resultMap = make(map[string]string)
	valid := token.ValidateToken(tokenString)
	if valid{
		resultMap["valid"] ="ok"
	}else{
		resultMap["valid"] ="fail"
	}
	c.Data["json"] = resultMap
	c.ServeJSON()
}

type TokenController struct {
	beego.Controller
}

func (c *TokenController) Get() {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	resultMap["token"] = token.NewToken()
	c.Data["json"] = resultMap
	c.ServeJSON()
}
