package main

import (
	_ "beego-filter/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"reflect"
	"beego-filter/controllers"
)

func main() {

	//===== 普通过滤,验证用户是否注册=========
	var allFilterUser = func(ctx *context.Context) {
		ctx.Output.Body([]byte("用户不存在,请登录"))
	}
	beego.InsertFilter("/*",beego.BeforeRouter,allFilterUser)

	//===== 正则路由进行过滤=========
	//为了方便测试,这里直接返回字符串
	//测试之前 先把 上面的 allFilterUser 注销 因为是拦截所有请求的
	//测试地址: http://localhost:8080/userlogin/55
	var FilterUser = func(ctx *context.Context) {
		ctx.Output.Body([]byte("用户不存在,请登录"))
	}
	beego.InsertFilter("/userlogin/:id([0-9]+)",beego.BeforeRouter,FilterUser)


	// 测试 过滤器实现路由
	//请求地址 ： http://localhost:8080/user
	//请求地址 ： http://localhost:8080/book
	var UrlManager = func(ctx *context.Context) {
		// 数据库读取全部的 url mapping 数据
		urlMapping := GetUrlMapping()
		for baseurl,rule:=range urlMapping {
			if baseurl == ctx.Request.RequestURI {
				ctx.Input.RunController = rule.controller
				ctx.Input.RunMethod = rule.method
				break
			}
		}
	}

	beego.InsertFilter("/*",beego.BeforeRouter,UrlManager)




	beego.Run()
}

func GetUrlMapping() map[string]rule {

	//这里是查询数据库  但是因为go不支持 根据字符串反射  所以需要自己查到映射 配置规则 例如 /user 开头的 都走 UserController
	urlMapping := make(map[string]rule)
	urlMapping["/user"] = rule{reflect.TypeOf(controllers.UserController{}),"Add"}
	urlMapping["/book"] = rule{reflect.TypeOf(controllers.BookController{}),"Del"}

	return urlMapping
}

type rule struct {
	//Controller名称
	controller reflect.Type
	//要执行的方法名称
	method string
}

