package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"fmt"
)

func init() {

	// bee 工具自动生成的固定路由
    beego.Router("/", &controllers.MainController{})

	/**
	基本Get路由实践
	注意事项：需要导入包：github.com/astaxie/beego/context
	测试用例：浏览器里输入：http://localhost:8080/get
	输出结果：在浏览器页面里显示：基本Get路由实践
	 */
	beego.Get("/get",func(ctx *context.Context){
		ctx.Output.Body([]byte("基本Get路由实践"))
	})

	/*
	基本Post路由实践
	注意事项：需要提交一个post请求
	输出结果：基本Post路由实践
	 */
	beego.Post("/post",func(ctx *context.Context){
		ctx.Output.Body([]byte("基本Post路由实践"))
	})

	/*
	响应任何 HTTP 的路由
	输出结果：响应任何 HTTP 的路由
	 */
	beego.Any("/any",func(ctx *context.Context){
		ctx.Output.Body([]byte("响应任何 HTTP 的路由"))
	})

	/**
	自定义的 handler 实现
	beego.Handler 第二个参数 就是你自己实现的 http.Handler
	http.Handler 下面有一个 ServeHTTP  你自定义一个结构体 实现该方法,那么就是 http.Handler了 。
	beego.Handler 第三个参数是否是前缀匹配,默认是 false, 如果设置了 true, 那么就会在路由匹配的时候前缀匹配,即 /rpc/user 这样的也会匹配去运行
	测试： 访问 http://localhost:8088/rpc   和  http://localhost:8088/rpc/user 都会 输出 TextHandler !
	 */
	thWelcome := &textHandler{"TextHandler !"}
	beego.Handler("/rpc", thWelcome,true)


	//======================正则路由开始========================

	/**
	默认匹配 //例如对于URL”/api/123”可以匹配成功，此时变量”:id”值为”123”
	测试用例：
	    能匹配：/api ; /api/123 ; /api/abc
	    不能匹配： /api/123/abc  即 /api/之后只能再接一个参数

	controller 中获取参数的代码
	func (c *RController) Get() {
		fmt.Println(c.Ctx.Input.Param(":id"));
	}
	 */
	beego.Router("/api/?:id", &controllers.RController{})

	/**
	默认匹配 //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123"，但URL"/api/"匹配失败
	测试用例：
	    能匹配： /api/123 ; /api/abc
	    不能匹配： /api/ ; /api/123/abc  即 /api/之后必须只能接受一个参数
	 */
	beego.Router("/api/:id", &controllers.RController{})

	/**
	自定义正则匹配 //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123"
	测试用例：
	    能匹配：/api/123 ; /api/345
	    不能匹配： /api/ ; /api/123/abc ; /api/123/123  即 /api/之后必须只能接受一个数字参数
	 */
	beego.Router("/api/:id([0-9]+)", &controllers.RController{})

	/**
	正则字符串匹配 //例如对于URL"/user/astaxie"可以匹配成功，此时变量":username"值为"astaxie"
	 */
	beego.Router("/user/:username([\\w]+)", &controllers.RController{})

	/**
	匹配方式 //例如对于URL"/download/file/api.xml"可以匹配成功，此时变量":path"值为"file/api"， ":ext"值为"xml"
	 */
	beego.Router("/download/*.*", &controllers.RController{})

	/**
	全匹配方式 //例如对于URL"/download/ceshi/file/api.json"可以匹配成功，此时变量":splat"值为"file/api.json"
	 */
	beego.Router("/download/ceshi/*", &controllers.RController{})

	/**
	int 类型设置方式，匹配 :id为int 类型，框架帮你实现了正则 ([0-9]+)
	 */
	beego.Router("/:id:int", &controllers.RController{})

	/*
	string 类型设置方式，匹配 :hi 为 string 类型。框架帮你实现了正则 ([\w]+)
	 */
	beego.Router("/:hi:string", &controllers.RController{})

	/**
	带有前缀的自定义正则 //匹配 :id 为正则类型。匹配 cms_123.html 这样的 url :id = 123
	 */
	beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{})

	//======================正则路由结束========================


	//======================自定义方法及 RESTful 规则开始======================

	//这里只测试两个,剩下的自行测试一下
	//所有的方法 路由到 IndexController 的 index 方法
	beego.Router("/Index",&controllers.IndexController{},"*:Index")

	//post请求 路由到 IndexController 的 CreateFood 方法
	beego.Router("/CreateFood",&controllers.IndexController{},"post:CreateFood")

	//是多个 HTTP Method 指向同一个函数的示例
	beego.Router("/ApiFunc",&controllers.RestController{},"get,post:ApiFunc")

	//测试不同的 method 对应不同的函数，通过 ; 进行分割  get 路由到 Index方法  post路由到 CreateFood 方法
	beego.Router("/simple",&controllers.IndexController{},"get:Index;post:CreateFood")

	//======================自定义方法及 RESTful 规则结束======================

	//自动匹配
	// http://localhost:8088/object/login  输出 {}
	// http://localhost:8088/object/login/name/wangsaichao/26 输出 ： {"0":"name","1":"wangsaichao","2":"26",":splat":"name/wangsaichao/26"}
	beego.AutoRouter(&controllers.ObjectController{})

	//注解路由 访问 http://localhost:8080/staticblock/aabb
	beego.Include(&controllers.CMSController{})

}

type textHandler struct {
	responseText string
}

func (th *textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, th.responseText)
}

