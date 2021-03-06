package models
import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)



type Master struct {
	Id                   int64
	Name                 string
	Official_url         string
	Product_description  string
	Invest_area          string
	Duration_description string
	Platform_est_rate    float32
	Platform_min_amount  float32
	Platform_icon_url    string
	Risk_rank            int64
}


func (this *Master) TableName() string {
	return "fin_p2p_master"
}


type MasterDao struct {

}

func (this *MasterDao) QueryByName(name string) (int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("fin_p2p_master")
	count, err := qs.Filter("name", name).Count()
	if err == nil {
		return count
	}else {
		beego.Error(err)
		return -1
	}
}


func (this *MasterDao) Query(name string) int64 {
	o := orm.NewOrm()
	m := new(Master)
	qs := o.QueryTable("fin_p2p_master")
	err := qs.Filter("name", name).One(m)
	if err==nil && m!=nil {
		return m.Id
	}else {
		return -1
	}
}


func (this *MasterDao) Save(m *Master) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	return id, err
}

func (this *MasterDao) SaveOrUpdate(m *Master) {
	o := orm.NewOrm()
	qs := o.QueryTable("fin_p2p_master")
	count, err := qs.Filter("name", m.Name).Count()
	if err ==nil && count==0 {
		o.Insert(m)
	}
	beego.Debug(count, err)
}