package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Redirect("/", 302)
}

func (this *ErrorController) ErrorParameterError() {
	fmt.Println("aaaaa")
	this.Data["json"] = "aaa"
	this.ServeJSON()
}
