package routers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"ssm/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/test", &controllers.TestController{})
	//beego.Router("/add", &controllers.AddController{})
	//beego.Router("/delete", &controllers.DeleteController{})
	ns := beego.NewNamespace("/docker",
		beego.NSBefore(auth),
		beego.NSRouter("/", &controllers.MainController{}),
		beego.NSRouter("/test", &controllers.TestController{}),
		beego.NSRouter("/add", &controllers.AddController{}),
		beego.NSRouter("/stop", &controllers.StopController{}),
		beego.NSRouter("/delete", &controllers.DeleteController{}))
	beego.AddNamespace(ns)
}

var auth = func(ctx *context.Context) {
	auth := ctx.Request.URL.Query().Get("auth")
	if len(auth) <= 0 || !controllers.AuthMaster(auth) {
		var ret = make(map[string]interface{})
		ret["status"] = false
		ret["data"] = "Are you kidding me?"
		//ctx.Redirect(302, "/login")
		b, _ := json.Marshal(ret)
		ctx.WriteString(string(b))
	}
}
