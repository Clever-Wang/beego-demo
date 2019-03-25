package main

import (
	_ "beego-build-url/routers"
	"github.com/astaxie/beego"
)

func main() {

	//fmt.Println(beego.URLFor("TestController.List"))
	//// 输出 /api/list
    //
	//fmt.Println(beego.URLFor("TestController.Get", ":last", "xie", ":first", "asta"))
	//// 输出 /person/xie/asta
    //
	//fmt.Println(beego.URLFor("TestController.Get", ":last", "xie", ":first", "asta"))
	//// 输出 /person/xie/asta
    //
	//fmt.Println(beego.URLFor("TestController.Params", "first", "wang", "second", "sai", "third", "chao"))
	//// 输出 /test/params?first=wang&second=sai&third=chao
    //
	//fmt.Println(beego.URLFor("TestController.Myext"))
	//// 输出 /Test/Myext
    //
	//fmt.Println(beego.URLFor("TestController.GetUrl"))
	//// 输出 /Test/GetUrl

	beego.URLFor("TestController.List")
	// 输出 /api/list

	beego.URLFor("TestController.Get", ":last", "xie", ":first", "asta")
	// 输出 /person/xie/asta

	beego.URLFor("TestController.Get", ":last", "xie", ":first", "asta")
	// 输出 /person/xie/asta

	beego.URLFor("TestController.Params", "first", "wang", "second", "sai", "third", "chao")
	// 输出 /test/params?first=wang&second=sai&third=chao

	beego.URLFor("TestController.Myext")
	// 输出 /Test/Myext

	beego.URLFor("TestController.GetUrl")
	// 输出 /Test/GetUrl


	beego.Run()
}

