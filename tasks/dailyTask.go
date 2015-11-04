package tasks

import (
	"time"
	"finprocess/crawlers"
)


func DailyTask() {

	go DailyReportInit(24*60) //每天一次

}


func DailyReportInit(minutes int) {

	timer := time.NewTicker(time.Minute * time.Duration(minutes))

	for {
		select {
		case <-timer.C:
			crawlers.BaiduCrawler()
			crawlers.Rong360Crawler()
			crawlers.Rong360bbsCrawler()
		}
	}
}
