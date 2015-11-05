package models

import (
	"github.com/astaxie/beego/orm"
)

//投资金额段 0:0-5万;1:5-10万;2:10-20万;3:20-50万;4:50万以上
//投资周期: 0:1个月以内,1:1-3个月;2:3-6个月;3:6个月-1年;4:1年以上
//预期收益, 0:0-5%;1:5-10%;2:10-15%;3:15%以上

const TableName = "fin_p2p_lov"

type Lov struct {
	Id   int64
	Code int
	Name string
	Type int
}

func (this *Lov) TableName() string {
	return TableName
}

type LovDao struct {
}

func (this *LovDao) Types() []Lov {
	var lovs []Lov
	o := orm.NewOrm()
	sql := "select * from fin_p2p_lov where type is null"
	count, err := o.Raw(sql).QueryRows(&lovs)
	if err == nil && count != 0 {
		return lovs
	} else {
		return nil
	}

}

func (this *LovDao) List(t int) []Lov {
	var lovs []Lov
	o := orm.NewOrm()
	count, err := o.QueryTable(TableName).Filter("type", t).All(&lovs)
	if err == nil && count != 0 {
		return lovs
	} else {
		return nil
	}
}
