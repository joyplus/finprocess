package main

import (
	"time"
	"finprocess/crawlers"
)

func main() {
	DailyReportInit(60)
}


func DailyReportInit(minutes int) {
	timer := time.NewTicker(time.Minute * time.Duration(minutes))

	for {
		select {
		case <-timer.C:
			crawlers.Crawler()
		}
	}
}
