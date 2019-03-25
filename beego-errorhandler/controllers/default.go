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

//测试错误处理
type ErrorHandlerController struct {
	beego.Controller
}

//访问直接跳转到401页面
func (this *ErrorHandlerController) To401() {
	this.Abort("401")
}

//main方法中自定义 404页面,测试跳转到自定义404
func (this *ErrorHandlerController) To404() {
	this.Abort("404")
}

//main方法中自定义 dbError错误,测试跳转到自定义dberror.tpl
func (this *ErrorHandlerController) Todberror() {
	this.Abort("dbError")
}

//测试跳转到自定义Controller 的501错误
func (this *ErrorHandlerController) To501() {
	this.Abort("501")
}


//自定义 ErrorController 在main方法中 要注册
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "501.tpl"
}

