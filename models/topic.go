package models
import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Topic struct {
	Id           int64 //need a primary key field
	Topic_id     int64
	Node_id      int64
	Uid          int64
	Ruid         int64
	Title        string
	Keywords     string
	Content      string
	Addtime      int64
	Updatetime   int64
	Lastreply    int
	Views        int
	Comments     int
	Favorites    int
	Closecomment int
	Is_top       int
	Is_hidden    int
	Ord          int
}



func (this *Topic) TableName() string {
	return "stb_topics"
}

type TopicDao struct {

}

func (this *TopicDao) InsertOrUpdate(t *Topic) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("stb_topics")
	count, err := qs.Filter("title", t.Title).Count()
	if err ==nil {
		if count==0 {
			o.Insert(t)
		}
	}else {
		beego.Error(err)
	}

	return 0
}