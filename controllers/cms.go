package controllers

import (
	"github.com/astaxie/beego"
)

type CmsController struct {
	beego.Controller
}

func (c *CmsController) Prepare() {
	// beego session使用
	c.SetSession("stat", int(1))
}

// @router /api/v1/cms:id [get]
func (c *CmsController) GetOne() {
	// beego session使用
	c.GetSession("stat")
	c.DelSession("stat")
	c.DestroySession()
}

// @router /api/v1/cms/user [post]
func (c *CmsController) Put() {
	// json格式数据
	c.Data["json"] = map[string]interface{}{"name": "astaxie"}
	c.ServeJSON()
	// 返回
	c.StopRun()
}