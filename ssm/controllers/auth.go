package controllers
import (
	"github.com/astaxie/beego"
)

func AuthMaster(master string) (result bool) {
	if master == beego.AppConfig.String("passauth") {
		return true
	}
	return false
}
