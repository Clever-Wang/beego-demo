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

//测试多种输出方式
type OutputController struct {
	beego.Controller
}

//直接输出字符串
func (c *OutputController) OutputString() {
	c.Ctx.WriteString("hello world")
}

//直接输出字符串第二种方式
func (c *OutputController) OutputString2() {
	c.Ctx.Output.Body([]byte("hello world"))
}

//模板数据输出
func (c *OutputController) OutputTpl() {
	c.TplName = "index.tpl"
}

//模板中放入数据
func (c *OutputController) OutputTplData() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//输出json内容
type JSONStruct struct {
	Code int
	Data interface{}
	Msg  string
}

func (c *OutputController) OutputJson() {
	jsonStruct := &JSONStruct{0, nil,"hello"}
	c.Data["json"] = jsonStruct
	c.ServeJSON()
}

//输出xml内容
type XMLStruct struct {
	Code int
	Data interface{}
	Msg  string
}

func (c *OutputController) OutputXml() {
	xmlStruct := &XMLStruct{0, "wangsaicaho","hello"}
	c.Data["xml"] = xmlStruct
	c.ServeXML()
}

//jsonp调用
func (c *OutputController) OutputJsonp() {
	jsonpStruct := &JSONStruct{0, "王赛超","hello"}
	c.Data["jsonp"] = jsonpStruct
	c.ServeJSONP()
}