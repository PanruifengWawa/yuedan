// @APIVersion 1.0.0
// @Title yuedan API
// @Description Document of APIs
// @Contact 46937922@qq.com
package routers

import (
	"yuedan/controllers"
	//	"yuedan/filters"

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
	//	beego.InsertFilter("/v1/*", beego.BeforeExec, filters.UserFilter)
	beego.AddNamespace(ns)
}
