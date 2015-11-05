package crawlers

import (
	"finprocess/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

func Rong360bbsCrawler() {

	for i := 1; i < 6; i++ {
		u := "http://bbs.rong360.com/forum-76-" + strconv.Itoa(i) + ".html"
		document, err := goquery.NewDocument(u)

		if err != nil {
			beego.Error(err)
			continue
		}

		//所有帖子
		document.Find("table#threadlisttableid").Find("tbody").Each(func(i int, selection *goquery.Selection) {
			topic := &models.Topic{}

			topic.Node_id = 4
			topic.Uid = 1

			t := selection.Find("th").First().Find("a.s.xst")

			title := t.Text()
			if len(title) > 0 {
				topic.Title = title

				if titleUrl, f := t.Attr("href"); f {
					//获取帖子正文
					c, err := goquery.NewDocument(titleUrl)

					if err == nil {
						content := c.Find("div#postlist").First().Find("td.t_f").First()

						content.Find("img").Each(func(i int, se *goquery.Selection) { // 替换图片的src地址
							if src, exists := se.Attr("file"); exists {
								se.SetAttr("src", "http://bbs.rong360.com/"+src)
							}
						})

						html, err := content.Html()
						if err == nil {
							topic.Content = html
							topic.Addtime = time.Now().Unix()
							topic.Updatetime = time.Now().Unix()

							(&models.TopicDao{}).InsertOrUpdate(topic)
						}
					}

					if err != nil {
						beego.Error(err)
					}

				}

			}
		})

	}

}
