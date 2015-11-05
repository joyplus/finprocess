package crawlers

import (
	"encoding/json"
	"finprocess/models"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func BaiduCrawler() {

	url := "http://caifu.baidu.com/wealth/ajax?pageSize=10&pageNum={pageNum}&module=Finance&category=wealth&serverTime=1445325063367&pvid=1445325063361173072&resourceid=1800181&subqid=1445325063361173072&sid=ui%3A0%26bsInsurance%3A3%26bsInvest%3A2%26bsLoan%3A1%26loanCardBbd%3A0&pssid=17522_1460_12772_14432_17001_17470_17073_15094_11723_16010_10633_17051&tn=baiduhome_pg&signTime=93&qid=14533663996696868932&wd=p2p&zt=ps&fr=ps&f=-&amount=&cycle=&profit=&productType=3102"
	for i := 1; i < 30; i++ {

		newUrl := strings.Replace(url, "{pageNum}", strconv.Itoa(i), -1)

		resp, err := http.Get(newUrl)
		if err != nil {
			beego.Error(err)
			continue
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error(err)
			continue
		}

		var dat map[string]interface{}

		e := json.Unmarshal(body, &dat)
		if e != nil {
			beego.Error(e)
			continue
		}

		if data, ok := dat["data"].(map[string]interface{}); ok {

			if li, ok := data["list"].([]interface{}); ok {

				if len(li) == 0 {
					continue
				}

				for _, v := range li {
					if product, ok := v.(map[string]interface{}); ok {

						//赋值
						master := &models.Master{}
						if supplierNameShort, ok := product["supplierNameShort"].(string); ok {
							master.Name = supplierNameShort
						}
						if supplierLogoUrl, ok := product["supplierLogoUrl"].(string); ok {
							master.Platform_icon_url = supplierLogoUrl
						}
						if rawUrl, ok := product["rawUrl"].(string); ok {
							master.Official_url = rawUrl
						}
						if idea, ok := product["idea"].(string); ok {
							master.Product_description = idea
						}

						//save master
						(&models.MasterDao{}).SaveOrUpdate(master)

						invest := &models.Invest_Contract{}
						invest.For_register = models.For_Register

						invest.Amount_max = 999999999
						if title, ok := product["title"].(string); ok {
							invest.Name = title
						}
						if idea, ok := product["idea"].(string); ok {
							invest.Description = idea
						}
						if lowestAmount, ok := product["lowestAmount"].(string); ok {
							if strings.Contains(lowestAmount, "元") {
								amount_min, err := strconv.Atoi(strings.Split(lowestAmount, "元")[0])
								if err == nil {
									invest.Amount_min = amount_min
								}
							} else if strings.Contains(lowestAmount, "万") {
								amount_min, err := strconv.Atoi(strings.Split(lowestAmount, "万")[0])
								if err == nil {
									invest.Amount_min = amount_min * 10000
								}
							}
						}
						invest.Duration_type = 2
						if investCycle, ok := product["investCycle"].(string); ok {
							if strings.Contains(investCycle, "个") {
								duration_min, err := strconv.Atoi(strings.Split(investCycle, "个")[0])
								if err == nil {
									invest.Duration_min = duration_min
								}
							}
						}
						if expectedProfitRate, ok := product["expectedProfitRate"].(string); ok {
							if rate, err := strconv.ParseFloat(expectedProfitRate, 32); err == nil {
								invest.Rate = rate
							}
						}

						if extraFields, ok := product["extraFields"].([]interface{}); ok && extraFields[0] != nil {
							if ex, ok1 := extraFields[0].([]interface{}); ok1 && ex[0] != nil {
								if m, ok2 := ex[0].(map[string]interface{}); ok2 {
									if tips, ok3 := m["tips"]; ok3 {
										switch tips {
										case "可提前赎回，可提前转让":
											invest.Early_terminate = 0
										case "可提前赎回，不可提前转让":
											invest.Early_terminate = 1
										case "不可提前赎回，可提前转让":
											invest.Early_terminate = 2
										case "不可提前赎回，不可提前转让":
											invest.Early_terminate = 3
										}
									}
								}
							}

						}

						//save invest_contract
						masterName := strings.Split(invest.Name, " - ")[0]
						masterId := (&models.MasterDao{}).Query(masterName)
						invest.Master_id = masterId
						(&models.Invest_ContractDao{}).SaveOrUpdate(invest)

					}
				}

			}

		}

	}
}
