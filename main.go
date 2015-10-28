package main

import (
	m "finprocess/models"
	_ "finprocess/docs"
	_ "finprocess/routers"

	"github.com/astaxie/beego"
)

func main() {
	m.Connect()
	beego.Run()
}
