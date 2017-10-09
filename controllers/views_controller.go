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
	c.TplName = "index.tpl"
}

func (c *RegistController) Get() {
	c.TplName = "regist.tpl"
}

func (c *FinishController) Get() {
	c.TplName = "finish.tpl"
}
func (c *MainController) Get() {
	c.TplName = "main.tpl"
}

func (c *MineController) Get() {
	c.TplName = "mine.tpl"
}

func (c *CenterController) Get() {
	c.TplName = "center.tpl"
}

func (c *CollectController) Get() {
	c.TplName = "collect.tpl"
}

func (c *ConfigController) Get() {
	c.TplName = "config.tpl"
}

func (c *PwdController) Get() {
	c.TplName = "pwd.tpl"
}

func (c *CreateAuctionController) Get() {
	c.TplName = "create_auction.tpl"
}
func (c *CreateOrderController) Get() {
	c.TplName = "create_order.tpl"
}

func (c *CreatePostController) Get() {
	c.TplName = "create_post.tpl"
}
