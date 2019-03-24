package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// 显示设置信息
func (c *MainController) Get() {
	flash:=beego.ReadFromRequest(&c.Controller)
	if _,ok:=flash.Data["notice"];ok{
		// 显示设置成功
		c.TplName = "set_success.tpl"
	}else if _,ok=flash.Data["error"];ok{
		// 显示错误
		c.TplName = "set_error.tpl"
	}else{
		// 不然默认显示设置页面
		c.Data["list"]= "跳转到默认设置页面"
		c.TplName = "setting_list.tpl"
	}
}

// 处理设置信息
func (c *MainController) Post() {
	flash := beego.NewFlash()
	flashType := c.GetString("flashType")
	if flashType == "1" {
		flash.Notice("1 代表设置成功")
		flash.Store(&c.Controller)
		c.Redirect("/",302)
		return
	}

	if flashType == "2" {
		flash.Error("2 代表设置失败")
		flash.Store(&c.Controller)
		c.Redirect("/",302)
		return
	}

	flash.Warning("3 代表显示默认")
	flash.Store(&c.Controller)
	c.Redirect("/",302)
}
