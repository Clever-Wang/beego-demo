package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
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

//======测试 开启 xsrf 攻击=============
type XSRFController struct {
	beego.Controller
}

//开启 xsrf 之后 使用Postmain 访问该接口 将无法访问
func (this *XSRFController) Post() {
	this.Ctx.Output.Body([]byte("测试开启 XSRF 攻击"))
}

//测试修改 xsrf 过期时间 需要在views 下面添加新的 xsrf.tpl页面
//1.先通过get请求跳转到 xsrf.tpl 页面
//2.访问 post 的方法,这个时候是可以访问的,访问之后  后退
//3.等待20秒后再次访问
func (this *XSRFController) Get(){
	this.XSRFExpire = 5 //测试发现无效，该值并没有用, 只有设置全局时间为 5s 之后再次请求才被禁止
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
	this.TplName = "xsrf.tpl"
}

//========测试jquery扩展================

type XSRFJqueryController struct {
	beego.Controller
}

//先访问get请求 跳转到 xsrf_js.tpl 页面
//点击页面上的 按钮 提交测试 alert  success
//过5秒再访问  没有反应 后台打印 Handler crashed with error XSRF cookie does not match POST argument
func (this *XSRFJqueryController) Get(){
	this.Data["xsrf_token"] = this.XSRFToken()
	this.TplName = "xsrf_js.tpl"
}

func (this *XSRFJqueryController) Test(){

	fmt.Println("========> 测试post 请求")
	this.Data["json"] = map[string]interface{}{"msg": "success"}
	this.ServeJSON()
}

