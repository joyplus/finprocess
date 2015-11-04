package models
import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
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

	Master          *Master        `orm:"-"` //不对Master对象orm映射,不然会有Master_id重复的错误
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
	if err ==nil {
		return count
	}else {
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
	var count int64
	var err error
	if len(i.Platform_id)==0 {
		count, err = qs.Filter("name", i.Name).Count()
	}else {
		count, err = qs.Filter("platform_id", i.Platform_id).Count()
	}
	if err ==nil && count==0 {
		o.Insert(i)
	}

	beego.Debug(count, err)

}


//查询所有产品,目前只显示每个平台rate最高的
func (this *Invest_ContractDao) ListAll() []Invest_Contract {
	o := orm.NewOrm()
	var invests []Invest_Contract

	sql := "select * from fin_p2p_invest_contract group by master_id order by rate desc"

	count, err := o.Raw(sql).QueryRows(&invests)
	if err ==nil && count!=0 {
		for i, in := range invests {
			if in.Master_id!=-1 {
				sql_master := "select * from fin_p2p_master where id = "+strconv.Itoa(int(in.Master_id))
				o1 := orm.NewOrm()
				m := &Master{}
				er := o1.Raw(sql_master).QueryRow(m)
				if er==nil {
					invests[i].Master=m
				}
			}
		}
		return invests
	}
	return nil
}
