package main

import (
	m "adexchange/models"
	_ "finprocess/routers"

	"finprocess/tasks"
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

	//minutes for duration
	duration, _ := beego.AppConfig.Int("task_duration")

	go tasks.DailyTaskInit(duration) //启动task

	beego.Run()

}
