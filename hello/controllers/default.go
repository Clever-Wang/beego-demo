package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

type MainController struct {
	beego.Controller
}

type RController struct {
	beego.Controller
}

type CmsController struct {
	beego.Controller
}

type IndexController struct {
	beego.Controller
}

type RestController struct {
	beego.Controller
}

type ObjectController struct {
	beego.Controller
}

type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
}

// @router /staticblock/:key [get]
func (c *CMSController) StaticBlock() {
	c.Ctx.Output.Body([]byte("测试注解路由: "+ c.Ctx.Input.Param(":key")))
}

func (c *ObjectController) Login() {

	param := c.Ctx.Input.Param(":ext")
	if param != ""{
		c.Ctx.Output.Body([]byte(param))
	}
	params := c.Ctx.Input.Params()
	if params != nil{
		str, _ := json.Marshal(params)
		c.Ctx.Output.Body([]byte(str))
	}
}

func (c *RestController) ApiFunc() {
	c.Ctx.Output.Body([]byte("多个 HTTP Method 指向同一个函数的示例"))
}

func (c *RController) Get() {
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":id")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":username")))
	c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":splat")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":path")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":ext")))
}

func (c *CmsController) Get() {
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":id")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":username")))
	c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":id")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":path")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":ext")))
}

func (c *IndexController) Index() {
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":id")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":username")))
	c.Ctx.Output.Body([]byte("index方法"))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":path")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":ext")))
}

func (c *IndexController) CreateFood() {
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":id")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":username")))
	c.Ctx.Output.Body([]byte("CreateFood方法"))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":path")))
	//c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":ext")))
}




type Config struct {
	AppName             string //Application name
	RunMode             string //Running Mode: dev | prod
	RouterCaseSensitive bool
}



func (c *MainController) Get() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)


	confighello := Config{
		AppName: "hello",
		RunMode: "aaa",
	}

	beego.Info(confighello)

	log.Info("string","fdsfadsfadsfdsafsadfdsa")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
