package main

import (
	m "adexchange/models"
	"finprocess/crawlers"
	_ "finprocess/routers"
	//"finprocess/tasks"
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

	//go tasks.DailyTaskInit(1) //启动task

	crawlers.Rong360Crawler()
	crawlers.Rong360bbsCrawler()

	beego.Run()

}
