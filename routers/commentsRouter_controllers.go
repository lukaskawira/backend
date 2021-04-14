package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["backend/controllers:ReservationController"] = append(beego.GlobalControllerRouter["backend/controllers:ReservationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:ReservationController"] = append(beego.GlobalControllerRouter["backend/controllers:ReservationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:rid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:ReservationController"] = append(beego.GlobalControllerRouter["backend/controllers:ReservationController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:rid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
