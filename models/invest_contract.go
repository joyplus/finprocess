package models

import (
	"fmt"
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

//时间类型的值,如果使用o.Insert o.Update方法,go会自动转成UTC时间保存到数据库
func (this *Invest_ContractDao) Save(i *Invest_Contract) {
	o := orm.NewOrm()
	sql := "INSERT INTO fin_p2p_invest_contract (platform_id, name, master_id, description, amount_min, amount_max, duration_min, duration_max, duration_type, for_register, rate, early_terminate, risk_rank, product_url, disp_order, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := o.Raw(sql, i.Platform_id, i.Name, i.Master_id, i.Description, i.Amount_min, i.Amount_max, i.Duration_min, i.Duration_max, i.Duration_type, i.For_register, i.Rate, i.Early_terminate, i.Risk_rank, i.Product_url, i.Disp_order, i.Create_time, i.Update_time).Exec()
	if err != nil {
		beego.Error(err)
	}
}

func (this *Invest_ContractDao) Update(i *Invest_Contract) {
	o := orm.NewOrm()
	sql := "UPDATE fin_p2p_invest_contract SET platform_id = ?, name = ?, master_id = ?, description = ?, amount_min = ?, amount_max = ?, duration_min = ?, duration_max = ?, duration_type = ?, for_register = ?, rate = ?, early_terminate = ?, risk_rank = ?, product_url = ?, disp_order = ?, create_time = ?, update_time = ? WHERE id = ?"
	_, err := o.Raw(sql, i.Platform_id, i.Name, i.Master_id, i.Description, i.Amount_min, i.Amount_max, i.Duration_min, i.Duration_max, i.Duration_type, i.For_register, i.Rate, i.Early_terminate, i.Risk_rank, i.Product_url, i.Disp_order, i.Create_time, i.Update_time, i.Id).Exec()
	if err != nil {
		beego.Error(err)
	}
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

	now, err := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	if err != nil {
		beego.Error(err)
	}

	if exist {
		i.Update_time = now
		this.Update(dao)
	} else {
		i.Create_time = now
		i.Update_time = now
		this.Save(i)
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

//执行rong360task时,删除任务执行前的数据
func (this *Invest_ContractDao) DeleteAllRong360(startTime time.Time) {
	o := orm.NewOrm()
	t := startTime.Format("2006-01-02 15:04:05")
	fmt.Println(t)
	sql := "delete from `fin_p2p_invest_contract` where LENGTH(platform_id)>0 and update_time < '" + t + "'"
	_, err := o.Raw(sql).Exec()
	if err != nil {
		beego.Error(err)
	}
}
