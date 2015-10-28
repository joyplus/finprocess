package models
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug=false

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3307)/p2p?charset=utf8")

	orm.RegisterModel(new(Invest_Contract))
	orm.RegisterModel(new(Lov))
	orm.RegisterModel(new(Master))
}