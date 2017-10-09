package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) RootView() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "index.tpl"

	c.Render()
}

func (c *HomeController) HomeView() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "index.tpl"
	c.Render()
}
