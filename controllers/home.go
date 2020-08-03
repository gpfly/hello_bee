package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c)
	c.TplName = "home.html"
}

func checkAccount(c *MainController) bool {

	ck, err := c.Ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = c.Ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	beego.Debug("uname:", uname, " pwd:", pwd)
	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd
}
