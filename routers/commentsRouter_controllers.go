package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["yuedan/controllers:TokenController"] = append(beego.GlobalControllerRouter["yuedan/controllers:TokenController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yuedan/controllers:TokenController"] = append(beego.GlobalControllerRouter["yuedan/controllers:TokenController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yuedan/controllers:TokenController"] = append(beego.GlobalControllerRouter["yuedan/controllers:TokenController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yuedan/controllers:TokenController"] = append(beego.GlobalControllerRouter["yuedan/controllers:TokenController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yuedan/controllers:UserController"] = append(beego.GlobalControllerRouter["yuedan/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yuedan/controllers:UserController"] = append(beego.GlobalControllerRouter["yuedan/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserDetails",
			Router: `/details`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yuedan/controllers:UserController"] = append(beego.GlobalControllerRouter["yuedan/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yuedan/controllers:UserController"] = append(beego.GlobalControllerRouter["yuedan/controllers:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
