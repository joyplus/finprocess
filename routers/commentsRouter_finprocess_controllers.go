package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["finprocess/controllers:InvestContractController"] = append(beego.GlobalControllerRouter["finprocess/controllers:InvestContractController"],
		beego.ControllerComments{
			"ListAll",
			`/list`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["finprocess/controllers:InvestContractController"] = append(beego.GlobalControllerRouter["finprocess/controllers:InvestContractController"],
		beego.ControllerComments{
			"ViewList",
			`/view`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["finprocess/controllers:LovController"] = append(beego.GlobalControllerRouter["finprocess/controllers:LovController"],
		beego.ControllerComments{
			"Types",
			`/types`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["finprocess/controllers:LovController"] = append(beego.GlobalControllerRouter["finprocess/controllers:LovController"],
		beego.ControllerComments{
			"List",
			`/list/:t`,
			[]string{"get"},
			nil})

}
