package controllers
import (
	"github.com/astaxie/beego"
	"finprocess/models"
)

type LovController struct {
	beego.Controller
}

// @Title 查询所有类型
// @Description 查询所有分类
// @router /types [get]
func (this *LovController) Types() {
	lovDao := &models.LovDao{}
	lovs := lovDao.Types()
	this.Data["json"]=lovs
	this.ServeJson()
}

// @Title 标签查询
// @Description 查询一个分类下的所有标签
// @Param type query int true "type value"
// @router /list/:t [get]
func (this *LovController) List() {
	t, err := this.GetInt(":t")
	if err==nil {
		lovDao := &models.LovDao{}
		lovs := lovDao.List(t)
		this.Data["json"]=lovs
	}
	this.ServeJson()
}
