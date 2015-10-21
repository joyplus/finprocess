package main

import (
	"time"
)

func main() {
	DailyReportInit(1)
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
