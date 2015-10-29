package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/lov","description":""},{"path":"/invest","description":""}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/invest":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/invest","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/list","description":"","operations":[{"httpMethod":"GET","nickname":"查询所有投资产品","type":"","summary":"只显示每个平台rate最高的产品"}]}]},"/lov":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/lov","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/types","description":"","operations":[{"httpMethod":"GET","nickname":"查询所有分类","type":"","summary":"查询所有分类"}]},{"path":"/list/:t","description":"","operations":[{"httpMethod":"GET","nickname":"标签查询","type":"","summary":"查询一个分类下的所有标签","parameters":[{"paramType":"query","name":"type","description":"\"type value\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	if beego.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocApi["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.Apis {
				a.Path = urlReplace(k + a.Path)
				v.Apis[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocApi[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
