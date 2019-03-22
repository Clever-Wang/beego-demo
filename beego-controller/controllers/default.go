package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//==================通过重写方法可以实现对应 method 的逻辑================
type AddController struct {
	beego.Controller
}

func (this *AddController) Prepare() {
	fmt.Println("=====> Prepare 函数执行了 ")
}

func (this *AddController) Get() {
	this.Ctx.Output.Body([]byte("=====> 这是一个Get请求"))
}

func (this *AddController) Post() {
	this.Ctx.Output.Body([]byte("=====> 这是一个Post请求"))
}


//==================比较流行的架构==================

type NestPreparer interface {
	NestPrepare()
}

// 定义一个BaseController 其他的Controller继承该 Controller 实现全局配置
type baseController struct {
	beego.Controller
}
// Prepare implemented Prepare method for baseRouter.
func (this *baseController) Prepare() {

    fmt.Println("=====>BaseController Prepare 函数执行了 ")

    //这里写每个 Controller 都有的逻辑

    //判断当前运行的 Controller 是否是 NestPreparer 实现
	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

type BaseAdminRouter struct {
	baseController
}

func (this *BaseAdminRouter) NestPrepare() {

    fmt.Println("=====>BaseAdminRouter NestPrepare 函数执行了 ")

    //这里写 BaseAdminRouter  特有的逻辑

}

func (this *BaseAdminRouter) Get(){
    this.Ctx.Output.Body([]byte("=====> 这是 BaseAdminRouter 的 Get 方法 "))
}

func (this *BaseAdminRouter) Post(){
    this.Ctx.Output.Body([]byte("=====> 这是 BaseAdminRouter 的 Post 方法 "))
}


//===============提前终止运行=================

type RController struct {
    beego.Controller
}

func (this *RController) Prepare() {
    this.Data["json"] = map[string]interface{}{"name": "astaxie"}
    this.ServeJSON()
    this.StopRun()
}

func (this *RController) Get() {
    this.Ctx.Output.Body([]byte("=====> 这是 测试提前终止的 的 Get 方法 "))
}
