// @APIVersion 1.0.0
// @Title yuedan API
// @Description Document of APIs
// @Contact 46937922@qq.com
package routers

import (
	"yuedan/controllers"
	"yuedan/filters"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{})
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/t_user",
			beego.NSInclude(
				&controllers.TUserController{},
			),
		),

		beego.NSNamespace("/t_token",
			beego.NSInclude(
				&controllers.TTokenController{},
			),
		),
	)
	beego.InsertFilter("/v1/*", beego.BeforeExec, filters.UserFilter)
	beego.AddNamespace(ns)
}
