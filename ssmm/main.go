package main

import (
	"os"
	"ssmm/controllers"
	_ "ssmm/routers"
	"ssmm/template"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	_ "github.com/go-sql-driver/mysql"
)

const (
	APP_VER = "0.1.0"
)

// We have to call a initialize function manully
// because we use `bee bale` to pack static resources
// and we cannot make sure that which init() execute first.
func initialize() {

	controllers.IsPro = beego.BConfig.RunMode == "prod"
	if controllers.IsPro {
		beego.SetLevel(beego.LevelInformational)
		os.Mkdir("./log", os.ModePerm)
		beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
	}

	controllers.InitApp()
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:Legend@GFW404@tcp(192.210.219.20:3306)/shadowsocks")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		// beego.BConfig.WebConfig.DirectoryIndex = true
		// beego.StaticDir["/swagger"] = "swagger"
	}
	initialize()
	if !controllers.IsPro {
		beego.SetStaticPath("/static_source", "static_source")
		beego.BConfig.WebConfig.DirectoryIndex = true
	}
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("minus", template.Minus)
	beego.Run()
}
