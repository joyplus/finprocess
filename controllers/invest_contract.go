package controllers
import (
	"github.com/astaxie/beego"
//	"finprocess/models"
	"finprocess/models"
)

type InvestContractController struct {
	beego.Controller
}

// @Title 查询所有投资产品
// @Description 只显示每个平台rate最高的产品
// @router /list [get]
func (this *InvestContractController) ListAll() {
	investDao := &models.Invest_ContractDao{}
	invests := investDao.ListAll()
	this.Data["json"]=invests
	this.ServeJson()
}