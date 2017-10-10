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
	beego.Router("/login", &controllers.HomeController{})
	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/finish", &controllers.FinishController{})
	beego.Router("/main", &controllers.MainController{})
	beego.Router("/center", &controllers.CenterController{})
	beego.Router("/mine", &controllers.MineController{})
	beego.Router("/collect", &controllers.CollectController{})
	beego.Router("/config", &controllers.ConfigController{})
	beego.Router("/pwd", &controllers.PwdController{})
	beego.Router("/create/order", &controllers.CreateOrderController{})
	beego.Router("/create/post", &controllers.CreatePostController{})
	beego.Router("/create/auction", &controllers.CreateAuctionController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/token",
			beego.NSInclude(
				&controllers.TokenController{},
			),
		),
	)
	beego.InsertFilter("/v1/user/update", beego.BeforeExec, filters.UserFilter)
	beego.InsertFilter("/v1/user/details", beego.BeforeExec, filters.UserFilter)
	beego.AddNamespace(ns)
}
