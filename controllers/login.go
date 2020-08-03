package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (m *LoginController) Get() {
	isExit := m.Input().Get("exit") == "true"
	if isExit {
		m.Ctx.SetCookie("uname", "", -1, "/")
		m.Ctx.SetCookie("pwd", "", -1, "/")
		m.Redirect("/", 301)
		return
	}
	m.TplName = "login.html"
}

func (m *LoginController) Post() {
	uname := m.Input().Get("uname")
	pwd := m.Input().Get("pwd")
	autoLogin := m.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		m.Ctx.SetCookie("uname", uname, maxAge, "/")
		m.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		m.Redirect("/", 301)
	}

	m.Redirect("/login", 301)
	return
}
