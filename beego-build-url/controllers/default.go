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

//====测试url构建====
type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	this.Data["Username"] = "astaxie"
	this.Ctx.Output.Body([]byte("ok"))
}

func (this *TestController) List() {
	this.Ctx.Output.Body([]byte("i am list"))
}

func (this *TestController) Params() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Params()[":first"] + this.Ctx.Input.Params()[":second"] + this.Ctx.Input.Params()[":third"]))
}

func (this *TestController) Myext() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
}

func (this *TestController) GetUrl() {
	this.Ctx.Output.Body([]byte(this.URLFor(".Myext")))
}
