package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["titd/controllers/controllers:TweetController"] = append(beego.GlobalControllerRouter["titd/controllers/controllers:TweetController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers/controllers:TweetController"] = append(beego.GlobalControllerRouter["titd/controllers/controllers:TweetController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers/controllers:TweetController"] = append(beego.GlobalControllerRouter["titd/controllers/controllers:TweetController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:tid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers/controllers:TwitterConfigController"] = append(beego.GlobalControllerRouter["titd/controllers/controllers:TwitterConfigController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers/controllers:TwitterConfigController"] = append(beego.GlobalControllerRouter["titd/controllers/controllers:TwitterConfigController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers/controllers:TwitterConfigController"] = append(beego.GlobalControllerRouter["titd/controllers/controllers:TwitterConfigController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:tid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers:TweetController"] = append(beego.GlobalControllerRouter["titd/controllers:TweetController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers:TweetController"] = append(beego.GlobalControllerRouter["titd/controllers:TweetController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers:TweetController"] = append(beego.GlobalControllerRouter["titd/controllers:TweetController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:tid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers:TwitterConfigController"] = append(beego.GlobalControllerRouter["titd/controllers:TwitterConfigController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers:TwitterConfigController"] = append(beego.GlobalControllerRouter["titd/controllers:TwitterConfigController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["titd/controllers:TwitterConfigController"] = append(beego.GlobalControllerRouter["titd/controllers:TwitterConfigController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:tid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
