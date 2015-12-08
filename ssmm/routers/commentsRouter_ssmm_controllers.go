package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["ssmm/controllers:ServerController"] = append(beego.GlobalControllerRouter["ssmm/controllers:ServerController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:ServerController"] = append(beego.GlobalControllerRouter["ssmm/controllers:ServerController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:ServerController"] = append(beego.GlobalControllerRouter["ssmm/controllers:ServerController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:ServerController"] = append(beego.GlobalControllerRouter["ssmm/controllers:ServerController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:ServerController"] = append(beego.GlobalControllerRouter["ssmm/controllers:ServerController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:UserController"] = append(beego.GlobalControllerRouter["ssmm/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:UserController"] = append(beego.GlobalControllerRouter["ssmm/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:UserController"] = append(beego.GlobalControllerRouter["ssmm/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:UserController"] = append(beego.GlobalControllerRouter["ssmm/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:UserController"] = append(beego.GlobalControllerRouter["ssmm/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:AccountController"] = append(beego.GlobalControllerRouter["ssmm/controllers:AccountController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:AccountController"] = append(beego.GlobalControllerRouter["ssmm/controllers:AccountController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:AccountController"] = append(beego.GlobalControllerRouter["ssmm/controllers:AccountController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:AccountController"] = append(beego.GlobalControllerRouter["ssmm/controllers:AccountController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:AccountController"] = append(beego.GlobalControllerRouter["ssmm/controllers:AccountController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:BillController"] = append(beego.GlobalControllerRouter["ssmm/controllers:BillController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:BillController"] = append(beego.GlobalControllerRouter["ssmm/controllers:BillController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:BillController"] = append(beego.GlobalControllerRouter["ssmm/controllers:BillController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:BillController"] = append(beego.GlobalControllerRouter["ssmm/controllers:BillController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ssmm/controllers:BillController"] = append(beego.GlobalControllerRouter["ssmm/controllers:BillController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
