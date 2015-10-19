package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"fmt"
)

type BaiduProcesser struct {
}

func (this *BaiduProcesser) Process(p *page.Page) {
	document := p.GetHtmlParser()

	//	content, e := document.Html()
	//	if e==nil {
	//		fmt.Println(content)
	//	}

	products := document.Find("li.list-item").Nodes

	for _, v := range products {

		p := goquery.NewDocumentFromNode(v)

		//		For_register  name中包含是否新人专享
		name := p.Find("div.item-title").Text()
		fmt.Println(name);

		//url
		url, ok := p.Find("a").First().Attr("href")
		if ok {
			fmt.Println(url)
		}

		//Amount_min
		amount_min := p.Find("span.field-col.field-col3").First().Text()
		fmt.Println(amount_min)

		//		Duration_min
		//		Duration_type
		//		Early_terminate
		duration := p.Find("span.field-col.field-col1.field-col-extend").First().Text()
		fmt.Println(duration)

		//		rate
		rate := p.Find("div.item-content-right").Find("span.returns").First().Text()
		fmt.Println(rate)

	}

}

func (this *BaiduProcesser) Finish() {
	fmt.Println("end .....")
}


func main() {
	spider.NewSpider(new(BaiduProcesser), "TaskName").
	AddUrl("http://caifu.baidu.com/wealth?wd=p2p&qid=14533663996696868932&pssid=17522_1460_12772_14432_17001_17470_17073_15094_11723_16010_10633_17051&sid=ui%3A0%26bsInsurance%3A2%26bsInvest%3A4%26bsLoan%3A2%26loanCardBbd%3A0&tn=baiduhome_pg&fr=ps&zt=ps&productType=3102&category=3102", "html").// start url, html is the responce type ("html" or "json")
	AddPipeline(pipeline.NewPipelineConsole()).// print result on screen
	SetThreadnum(3).
	Run()
}