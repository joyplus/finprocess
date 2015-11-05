package tasks

import (
	"finprocess/crawlers"
	"github.com/astaxie/beego"
	"time"
)

//func DailyTask() {

//	go DailyTaskInit(1) //每天一次

//}

func DailyTaskInit(minutes int) {

	beego.Info("Start Daily Task")

	timer := time.NewTicker(time.Minute * time.Duration(minutes))

	for {
		select {
		case <-timer.C:
			//crawlers.BaiduCrawler()
			crawlers.Rong360Crawler()
			crawlers.Rong360bbsCrawler()
		}
	}
}
