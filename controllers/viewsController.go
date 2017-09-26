package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}
type RegistController struct {
	beego.Controller
}
type FinishController struct {
	beego.Controller
}
type MainController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Render()
}

func (c *RegistController) Get() {
	c.TplName = "regist.tpl"
	c.Render()
}

func (c *FinishController) Get() {
	c.TplName = "finish.tpl"
	c.Render()
}

func (c *MainController) Get() {
	c.TplName = "main.tpl"
	c.Render()
}
