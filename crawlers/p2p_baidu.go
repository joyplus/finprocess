package crawlers

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"finprocess/models"
	"github.com/astaxie/beego/orm"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"encoding/json"
)


func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3307)/p2p?charset=utf8")
	orm.RegisterModel(new(models.Invest_Contract), new(models.Master))

}


func Crawler() {

	url := "http://caifu.baidu.com/wealth/ajax?pageSize=10&pageNum={pageNum}&module=Finance&category=wealth&serverTime=1445325063367&pvid=1445325063361173072&resourceid=1800181&subqid=1445325063361173072&sid=ui%3A0%26bsInsurance%3A3%26bsInvest%3A2%26bsLoan%3A1%26loanCardBbd%3A0&pssid=17522_1460_12772_14432_17001_17470_17073_15094_11723_16010_10633_17051&tn=baiduhome_pg&signTime=93&qid=14533663996696868932&wd=p2p&zt=ps&fr=ps&f=-&amount=&cycle=&profit=&productType=3102"
	for i := 1; i<30; i++ {

		newUrl := strings.Replace(url, "{pageNum}", strconv.Itoa(i), -1)
		fmt.Println(newUrl)

		resp, err := http.Get(newUrl)
		if err!=nil {
			fmt.Println(err)
			continue
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err!=nil {
			fmt.Println(err)
			continue
		}



		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err==nil {

			if data, ok := dat["data"].(map[string]interface{}); ok {

				if li, ok := data["list"].([]interface{}); ok {

					if len(li)==0 {
						continue
					}

					for _, v := range li {
						if product, ok := v.(map[string]interface{}); ok {

							//赋值
							master := &models.Master{}
							if supplierNameShort, ok := product["supplierNameShort"].(string); ok {
								master.Name=supplierNameShort
							}
							if supplierLogoUrl, ok := product["supplierLogoUrl"].(string); ok {
								master.Platform_icon_url=supplierLogoUrl
							}
							if rawUrl, ok := product["rawUrl"].(string); ok {
								master.Official_url=rawUrl
							}
							if idea, ok := product["idea"].(string); ok {
								master.Product_description=idea
							}

							//save master
							(&models.MasterDao{}).SaveOrUpdate(master)

							invest_contract := &models.Invest_Contract{}
							invest_contract.For_register=1
							if title, ok := product["title"].(string); ok {
								invest_contract.Name=title
							}
							if idea, ok := product["idea"].(string); ok {
								invest_contract.Description=idea
							}
							if lowestAmount, ok := product["lowestAmount"].(string); ok {
								invest_contract.Amount_min=lowestAmount
							}
							if investCycle, ok := product["investCycle"].(string); ok {
								invest_contract.Duration_min=investCycle
							}
							if expectedProfitRate, ok := product["expectedProfitRate"].(string); ok {
								if rate, err := strconv.ParseFloat(expectedProfitRate, 32); err==nil {
									invest_contract.Rate=float32(rate)
								}
							}

							//fmt.Println(invest_contract)
							//fmt.Println(master)

							//save invest_contract
							(&models.Invest_ContractDao{}).SaveOrUpdate(invest_contract)

						}
					}
				}

			}

		}else {
			fmt.Println(err)
		}

	}
}
