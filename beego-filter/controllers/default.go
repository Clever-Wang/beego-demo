package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


//测试在 过滤器中 url 映射
type UserController struct {
	beego.Controller
}
func (this *UserController) Add() {
	this.Ctx.Output.Body([]byte("=====> 添加user成功"))
}


type BookController struct {
	beego.Controller
}

func (this *BookController) Del() {
	this.Ctx.Output.Body([]byte("=====> 删除book成功"))
}