package tests

import (
	"finprocess/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3307)/p2p?charset=utf8")
}

func TestP2p(t *testing.T) {

	//	crawlers.BaiduCrawler()

	//	crawlers.Rong360Crawler()

	//	crawlers.Rong360bbsCrawler()

}

func TestDelete(t *testing.T) {
	investDao := &models.Invest_ContractDao{}
	investDao.DeleteAllRong360()
}
