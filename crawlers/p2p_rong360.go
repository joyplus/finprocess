package crawlers
import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"finprocess/models"
	"net/http"
	"github.com/henrylee2cn/mahonia"
)

var client = &http.Client{}

func Rong360Crawler() {

	//只获取前30页的,防止被屏蔽
	for i := 1; i<30; i++ {

		u := "http://www.rong360.com/licai-p2p/list/p"+strconv.Itoa(i)
		document, _ := getDocument(u)

		document.Find("div.floor.clearfix").Find("table.wd-pro-list").Find("tbody").Find("tr").Each(func(i int, selection *goquery.Selection) {

			if i%2==0 {
				if clickurl, f := selection.Attr("click-url"); f {

					invest := &models.Invest_Contract{}

					//platform_id
					us := strings.Split(clickurl, "-")
					if len(us)>0 {
						platform_id := us[len(us)-1]
						invest.Platform_id=platform_id
					}

					var description string
					if s := selection.Find("td.td7"); s!=nil {
						description=strings.TrimSpace(s.Text())
					}

					if product, err := getDocument(clickurl); err==nil {

						master := &models.Master{}

						masterName := product.Find("div.t.clearfix").Find("div.r").Find("p").First().Text()
						master.Name=masterName

						if icon_url, exists := product.Find("div.t.clearfix").Find("div.l").Find("img").Attr("src"); exists {
							master.Platform_icon_url=icon_url
						}

						if u, exists := product.Find("div.view-r-des").Find("a.btn").Attr("href"); exists {
							if website, err := getDocument(u); err==nil {
								if u, f := website.Find("div.m.wrap-left").Find("div.p1").Find("a").Attr("href"); f {
									master.Official_url=u
								}
							}
						}

						//save master
						(&models.MasterDao{}).SaveOrUpdate(master)


						//title description
						invest.Description=description
						invest.For_register=models.For_Register
						invest.Amount_max=999999999
						title := product.Find("div.t").Find("span").Text()
						invest.Name=title

						//rate and duration
						rateAndDuration := product.Find("div.d").Find("p").Nodes
						if (len(rateAndDuration)==3) {
							rate_text := goquery.NewDocumentFromNode(rateAndDuration[0]).Find("span").Text()
							if rate, err := strconv.ParseFloat(strings.TrimRight(rate_text, "%"), 16); err==nil {
								invest.Rate=rate
							}

							durationNode := goquery.NewDocumentFromNode(rateAndDuration[1])
							if duration, err := strconv.Atoi(durationNode.Find("span").Text()); err==nil {
								invest.Duration_min=duration
							}

							durationTypeText := durationNode.Text()
							if strings.Contains(durationTypeText, "天") {
								invest.Duration_type=models.Duration_Type_Day
							} else if strings.Contains(durationTypeText, "月") {
								invest.Duration_type=models.Duration_Type_Month
							}else if strings.Contains(durationTypeText, "年") {
								invest.Duration_type=models.Duration_Type_Year
							}
						}

						clearfix := product.Find("div.c.clearfix").Find("table.tab1").Find("tr").Nodes
						if len(clearfix)==6 {
							amountNode := goquery.NewDocumentFromNode(clearfix[3])
							if amount, err := strconv.Atoi(strings.TrimLeft(strings.TrimRight(amountNode.Find(".td2").Text(), "万元"), "元")); err==nil {
								invest.Amount_min=amount
							}
						}

						//save invest_contract
						masterId := (&models.MasterDao{}).Query(masterName)
						invest.Master_id= masterId
						(&models.Invest_ContractDao{}).SaveOrUpdate(invest)

					}
				}
			}

		})
	}

}


func getDocument(url string) (*goquery.Document, error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	//	req.Header.Add("Accept-Encoding", "gzip, deflate, sdch")  //有这个会乱码
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", "www.rong360.com")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2490.71 Safari/537.36")

	resp, _ := client.Do(req)

	decoder := mahonia.NewDecoder("UTF-8")

	document, err := goquery.NewDocumentFromReader(decoder.NewReader(resp.Body))

	return document, err
}