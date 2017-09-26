// @APIVersion 1.0.0
// @Title yuedan API
// @Description Document of APIs
// @Contact 46937922@qq.com
package routers

import (
	"yuedan/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/finish", &controllers.FinishController{})
	beego.Router("/main", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/t_user",
			beego.NSInclude(
				&controllers.TUserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
