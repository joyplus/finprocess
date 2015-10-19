package models

import (
	"github.com/astaxie/beego/orm"
)

type Master struct {
	Id                   int       	   `orm:pk;auto`
	Name                 string        `orm:size(50)`
	Official_url         string        `orm:size(500)`
	Product_description  string        `orm:size(100)`
	Invest_area          string        `orm:size(50)`
	Duration_description string        `orm:size(100)`
	Platform_est_rate    float32
	Platform_min_amount  float32
	Platform_icon_url    string        `orm:size(10)`
	Risk_rank            int
}

type MasterDAO struct {
}

func (this *MasterDAO) Save(c *Master) (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(c)
	return id, err
}

func (this *MasterDAO) Get(id int) (m *Master) {
	o := orm.NewOrm()
	master := &Master{}
	qs := o.QueryTable("fin_p2p_master")
	qs.Filter("id", id).One(master)
	return master
}
