package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["bapi/controllers:AccountsController"] = append(beego.GlobalControllerRouter["bapi/controllers:AccountsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})


}
