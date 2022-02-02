// Package routers
// @APIVersion 1.0.0
// @Title tit api daemon
// @Description this api is core of the tit app, and works locally while tit cli talks to it via a client/http
// @Contact mail@karngyan.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"titd/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/tweets",
			beego.NSInclude(
				&controllers.TweetController{},
			),
		),
		beego.NSNamespace("/tconfigs",
			beego.NSInclude(
				&controllers.TwitterConfigController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
