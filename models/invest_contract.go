package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Invest_Contract struct {
	Id              int64
	Platform_id     string
	Name            string
	Master_id       int64
	Description     string
	Amount_min      int
	Amount_max      int
	Duration_min    int
	Duration_max    int
	Duration_type   int
	For_register    int
	Rate            float64
	Early_terminate int
	Risk_rank       int

	Master *Master `orm:"-"` //不对Master对象orm映射,不然会有Master_id重复的错误

	Product_url string
	Disp_order  int
	Create_time time.Time `orm:type(TIMESTAMP)`
	Update_time time.Time `orm:type(TIMESTAMP)`
}

func (this *Invest_Contract) TableName() string {
	return "fin_p2p_invest_contract"
}

type Invest_ContractDao struct {
}

func (this *Invest_ContractDao) QueryByName(name string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("fin_p2p_invest_contract")
	count, err := qs.Filter("name", name).Count()
	if err == nil {
		return count
	} else {
		beego.Error(err)
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

	//百度平台根据名字来判断是否存在,rong360根据platform_id
	var exist bool
	dao := &Invest_Contract{}

	if len(i.Platform_id) == 0 {
		exist = qs.Filter("name", i.Name).Exist()
		if exist {
			qs.Filter("name", i.Name).One(dao)
		}
	} else {
		exist = qs.Filter("platform_id", i.Platform_id).Exist()
		if exist {
			qs.Filter("platform_id", i.Platform_id).One(dao)
		}
	}

	now, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		now = time.Now().Local()
	} else {
		beego.Error(err)
	}

	if exist {
		o.Delete(dao)
		create, e := time.Parse("2006-01-02 15:04:05", dao.Create_time.Format("2006-01-02 15:04:05"))
		if e != nil {
			create = time.Now().Local()
		} else {
			beego.Error(e)
		}
		i.Create_time = create
		i.Update_time = now
		o.Insert(i)
	} else {
		i.Create_time = now
		i.Update_time = now
		o.Insert(i)
	}

}

//查询所有产品,目前只显示每个平台rate最高的
func (this *Invest_ContractDao) ListAll() []Invest_Contract {
	o := orm.NewOrm()
	var invests []Invest_Contract

	sql := "select * from fin_p2p_invest_contract group by master_id order by rate desc,disp_order desc,update_time desc"

	count, err := o.Raw(sql).QueryRows(&invests)
	if err == nil && count != 0 {
		for i, in := range invests {
			if in.Master_id != -1 {
				sql_master := "select * from fin_p2p_master where id = " + strconv.Itoa(int(in.Master_id))
				o1 := orm.NewOrm()
				m := &Master{}
				er := o1.Raw(sql_master).QueryRow(m)
				if er == nil {
					invests[i].Master = m
				}
			}
		}
		return invests
	}
	return nil
}

//执行rong360task时,删除10分钟前的数据,也就是上次更新的数据,即每次只保留新的30页数据
func (this *Invest_ContractDao) DeleteAllRong360() {
	o := orm.NewOrm()
	before := time.Now().Add(-10 * time.Minute).Format("2006-01-02 15:04:05")

	sql := "delete from `fin_p2p_invest_contract` where platform_id !=0 and update_time <= '" + before + "'"
	_, err := o.Raw(sql).Exec()

	if err != nil {
		beego.Error(err)
	}
}
