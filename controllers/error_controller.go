package controllers

import (
	"yuedan/enums"

	"yuedan/utils"

	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.TplName = "404.tpl"
}

func (this *ErrorController) ErrorParameterError() {
	this.Data["json"] = utils.GenerateDataWrapper(enums.ParameterError, "参数错误")
	this.ServeJSON()
}
