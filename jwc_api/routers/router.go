// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/scbizu/Zafu_jwcInterface.git/jwc_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		// beego.NSNamespace("/jwc",
		// 	beego.NSInclude(
		// 		&controllers.JwcController{},
		// 	),
		// ),
	)
	beego.AddNamespace(ns)
	beego.Router("/jwc", &controllers.JwcController{})
	beego.Router("/vrcode", &controllers.JwcVrController{})
	beego.Router("/test", &controllers.TestController{})
	//	beego.Router("/course", &controllers.CourseController{})
	beego.Router("/exam", &controllers.ExamController{})
	beego.Router("/score", &controllers.ScoreController{})
}
