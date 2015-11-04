package models
import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func init() {
	orm.Debug, _ = beego.AppConfig.Bool("orm_debug")

	orm.RegisterModel(new(Invest_Contract))
	orm.RegisterModel(new(Lov))
	orm.RegisterModel(new(Master))
	orm.RegisterModel(new(Topic))
}
