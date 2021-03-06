package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	Ord          int64
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
	if err == nil && count == 0 {
		o.Insert(t)
	}
	beego.Debug(count, err)

	return 0
}
