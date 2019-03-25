package main

import (
	_ "beego-errorhandler/routers"
	"github.com/astaxie/beego"
	"net/http"
	"html/template"
	"beego-errorhandler/controllers"
)
//重新定义 404 页面 ,其他页面类似
func pageNotFound(rw http.ResponseWriter, r *http.Request){
	t,_:= template.New("404.tpl").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/404.tpl")
	data :=make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

//自定义数据库错误
func dbError(rw http.ResponseWriter, r *http.Request){
	t,_:= template.New("dberror.tpl").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/dberror.tpl")
	data :=make(map[string]interface{})
	data["content"] = "database is now down"
	t.Execute(rw, data)
}

func main() {
	beego.ErrorHandler("404",pageNotFound)
	beego.ErrorHandler("dbError",dbError)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}


