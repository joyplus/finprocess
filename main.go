package main

import (
	_ "finprocess/docs"
	m "finprocess/models"
	_ "finprocess/routers"

	"github.com/astaxie/beego"
)

func main() {
	m.Connect()
	beego.ViewsPath = "views"
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"

	beego.AddTemplateExt("html")
	beego.SetStaticPath("/public", "public")

	beego.SetLogger("file", `{"filename":"logs/finprocess.log"}`)
	beego.SetLogFuncCall(true)

	logLevel, _ := beego.AppConfig.Int("log_level")
	beego.SetLevel(logLevel)
	beego.Run()
}
