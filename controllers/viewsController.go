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

type CenterController struct {
	beego.Controller
}
type MineController struct {
	beego.Controller
}
type CollectController struct {
	beego.Controller
}
type ConfigController struct {
	beego.Controller
}
type PwdController struct {
	beego.Controller
}
type Page404Controller struct {
	beego.Controller
}
type CreateOrderController struct {
	beego.Controller
}
type CreatePostController struct {
	beego.Controller
}
type CreateAuctionController struct {
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

func (c *MineController) Get() {
	c.TplName = "mine.tpl"
	c.Render()
}

func (c *CenterController) Get() {
	c.TplName = "center.tpl"
	c.Render()
}

func (c *CollectController) Get() {
	c.TplName = "collect.tpl"
	c.Render()
}

func (c *ConfigController) Get() {
	c.TplName = "config.tpl"
	c.Render()
}

func (c *Page404Controller) Get() {
	c.TplName = "404.tpl"
	c.Render()
}

func (c *PwdController) Get() {
	c.TplName = "pwd.tpl"
	c.Render()
}

func (c *CreateAuctionController) Get() {
	c.TplName = "create_auction.tpl"
	c.Render()
}
func (c *CreateOrderController) Get() {
	c.TplName = "create_order.tpl"
	c.Render()
}

func (c *CreatePostController) Get() {
	c.TplName = "create_post.tpl"
	c.Render()
}
