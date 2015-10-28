package models
import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
)


type Invest_Contract struct {
	Id              int64
	Name            string
	Master_id       int64
	Description     string
	Amount_min      int
	Amount_max      int
	Duration_min    int
	Duration_max    int
	Duration_type   int
	For_register    int
	Rate            float32
	Early_terminate int
	Risk_rank       int

	Master          *Master `orm:"-"` //Master  *Master这么写会导致循环依赖
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


//查询所有产品,目前只显示每个平台rate最高的
func (this *Invest_ContractDao) ListAll() []Invest_Contract {
	o := orm.NewOrm()
	var invests []Invest_Contract
	//	sql := "select * from ( select * from fin_p2p_invest_contract i group by i.master_id order by i.rate desc) invest left join fin_p2p_master master on master.id=invest.master_id"

	sql := "select * from fin_p2p_invest_contract group by master_id order by rate desc"

	count, err := o.Raw(sql).QueryRows(&invests)
	if err ==nil && count!=0 {
		for i, in := range invests {
			sql_master := "select * from fin_p2p_master where id = "+strconv.Itoa(int(in.Master_id))
			o1 := orm.NewOrm()
			m := &Master{}
			er := o1.Raw(sql_master).QueryRow(m)
			fmt.Println(i,in.Master_id,m)
			if er==nil {
				invests[i].Master=m
			}
		}
		return invests
	}
	return nil
}
