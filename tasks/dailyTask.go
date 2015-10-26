package tasks

import (
	"time"
	"finprocess/crawlers"
)

func dd() {
	DailyReportInit(1)
}


func DailyReportInit(minutes int) {
	timer := time.NewTicker(time.Second * time.Duration(minutes))

	for {
		select {
		case <-timer.C:
			crawlers.Crawler()
		}
	}
}
