// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"ssmm/controllers"
)

func init() {
	/*
		ns := beego.NewNamespace("/v1",

			beego.NSNamespace("/account",
				beego.NSInclude(
					&controllers.AccountController{},
				),
			),

			beego.NSNamespace("/bill",
				beego.NSInclude(
					&controllers.BillController{},
				),
			),

			beego.NSNamespace("/server",
				beego.NSInclude(
					&controllers.ServerController{},
				),
			),

			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
		)
		ns1 := beego.NewNamespace("api",

			beego.NSRouter("/server", &controllers.ServerController{}, "post:Post"),
		)
		beego.AddNamespace(ns)
		beego.AddNamespace(ns1)
	*/
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/service", &controllers.ServiceController{}, "get:GetActive")
	beego.Router("/cart/conf/:id([0-9]+)", &controllers.CartController{}, "get:ConfService")
	beego.Router("/cart/view", &controllers.CartController{}, "get:CheckoutService")
	beego.Router("/cart", &controllers.CartController{}, "get:ViewService")
	beego.Router("/cart/add", &controllers.CartController{}, "post:AddService")
	beego.Router("/cart/checkout", &controllers.CartController{}, "post:PlaceOrder")
	beego.Router("/user/register", &controllers.WebUserController{}, "get:GetRegister;post:Register")
	beego.Router("/user/login", &controllers.WebUserController{}, "get:GetLogin;post:Login")
	beego.Router("/user", &controllers.WebUserController{}, "get:GetHome")
	beego.Router("/pay", &controllers.PayController{}, "get:BeforePay;post:Pay")
	beego.Router("/pay/callback", &controllers.PayController{}, "get,post:Callback")
	beego.Router("/pay/result/:id([0-9]+)", &controllers.PayController{}, "get:PayResult")
}
