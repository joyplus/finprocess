package main

import (
	m "adexchange/models"
	_ "finprocess/docs"
	_ "finprocess/routers"

	"finprocess/tasks"
	"github.com/astaxie/beego"
)

func main() {

	m.Connect()

	tasks.DailyTask() //启动task

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
