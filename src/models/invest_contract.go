package models

import (
	"github.com/astaxie/beego/orm"
)

type Invest_Contract struct {
	Id              int            `orm:pk;auto`
	Master_id       int
	Description     string        `orm:size(100)`
	Amount_min      int
	Amount_max      int
	Duration_min    int
	Duration_max    int
	Duration_type   int
	For_register    int
	Rate            float32
	Early_terminate int
	Risk_rank       int
}

type Invest_ContractDAO struct {
}

func (this *Invest_ContractDAO) Save(i *Invest_Contract) (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(i)
	return id, err
}

func (this *Invest_ContractDAO) Get(id int) (m *Invest_Contract) {
	o := orm.NewOrm()
	i := &Invest_Contract{}
	qs := o.QueryTable("fin_p2p_master")
	qs.Filter("id", id).One(i)
	return i
}
