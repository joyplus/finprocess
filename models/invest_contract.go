package models
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)


type Invest_Contract struct {
	Id              int64
	Name            string
	Master_id       int64
	Description     string
	Amount_min      string
	Amount_max      string
	Duration_min    string
	Duration_max    string
	Duration_type   int64
	For_register    int64
	Rate            float32
	Early_terminate int64
	Risk_rank       int64
}


func (this *Invest_Contract) TableName() string {
	return "fin_p2p_invest_contract"
}

type Invest_ContractDao struct {

}

func (this *Invest_ContractDao) QueryByName(name string) (int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("fin_p2p_invest_contract")
	count, err := qs.Filter("name", name).Count()
	if err ==nil {
		return count
	}else {
		fmt.Println(err)
		return -1
	}
}

func (this *Invest_ContractDao) Save(i *Invest_Contract) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(i)
	return id, err
}

func (this *Invest_ContractDao) SaveOrUpdate(i *Invest_Contract) {
	o := orm.NewOrm()
	qs := o.QueryTable("fin_p2p_invest_contract")
	count, err := qs.Filter("name", i.Name).Count()
	if err ==nil {
		if count==0 {
			id, err := o.Insert(i)
			fmt.Println(id, err)
		}
	}else {
		fmt.Println(err)
	}
}
