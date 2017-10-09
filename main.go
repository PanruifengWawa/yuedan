package main

import (
	_ "yuedan/routers"

	"yuedan/controllers"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {

	//db parameters
	user := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlhost")
	port := beego.AppConfig.String("mysqlport")
	dbName := beego.AppConfig.String("mysqldb")
	mysqlurl := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&loc=Local"

	orm.RegisterDataBase("default", "mysql", mysqlurl)
	if beego.BConfig.RunMode == "dev" {
		//api documents
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		orm.Debug = true
	} else {
		orm.Debug = false
	}

	//error handler
	beego.ErrorController(&controllers.ErrorController{})

	//log
	beego.SetLogger("file", `{"filename":"logs/log.log"}`)

}

func main() {

	beego.Run()
}
