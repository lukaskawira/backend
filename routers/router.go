package routers

import (
	"backend/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/res",
			beego.NSInclude(
				&controllers.ReservationController{},
			),
		),
		beego.NSNamespace("/cus",
			beego.NSInclude(
				&controllers.CustomerController{},
			),
		),
		beego.NSNamespace("/con",
			beego.NSInclude(
				&controllers.ContactController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{})
}
