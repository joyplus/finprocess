package main
import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"time"
	"finprocess/models"
)


func main() {

	for i := 1; i<6; i++ {
		u := "http://bbs.rong360.com/forum-76-"+strconv.Itoa(i)+".html"
		document, _ := goquery.NewDocument(u)

		nodes := document.Find("table#threadlisttableid").Find("tbody").Nodes

		//所有帖子
		for _, n := range nodes {
			topic := &models.Topic{}

			doc := goquery.NewDocumentFromNode(n)
			t := doc.Find("th").First().Find("a.s.xst")

			title := t.Text()
			if title==nil || len(title)==0 {
				continue
			}
			topic.Title=title

			titleUrl, flag := t.Attr("href")
			if flag {
				//获取帖子正文
				c, _ := goquery.NewDocument(titleUrl)
				content, _ := c.Find("div#postlist").First().Find("td.t_f").First().Html()
				topic.Content=content

			}
			topic.Addtime=time.Now().Unix()
			topic.Updatetime=time.Now().Unix()

			(&models.TopicDao{}).InsertOrUpdate(topic)
		}
	}


	//title keywords contents
}
