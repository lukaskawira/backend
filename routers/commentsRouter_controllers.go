package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["backend/controllers:ContactController"] = append(beego.GlobalControllerRouter["backend/controllers:ContactController"],
        beego.ControllerComments{
            Method: "Insert",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:cid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:cid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:CustomerController"] = append(beego.GlobalControllerRouter["backend/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout/:cid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
            Method: "Cancel",
            Router: "/:rid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:ReservationController"] = append(beego.GlobalControllerRouter["backend/controllers:ReservationController"],
        beego.ControllerComments{
            Method: "GetResByCustomerID",
            Router: "/gcust/:cid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend/controllers:ReservationController"] = append(beego.GlobalControllerRouter["backend/controllers:ReservationController"],
        beego.ControllerComments{
            Method: "RealtimeGet",
            Router: "/realt/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
