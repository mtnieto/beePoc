// @APIVersion 1.0.0
// @Title Series API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact mtnieto94@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/beePoc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/series",
			beego.NSInclude(
				&controllers.SerieController{},
			),
		),
		
	)
	beego.AddNamespace(ns)
}
