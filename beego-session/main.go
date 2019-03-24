package main

import (
	_ "beego-session/routers"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "root:123456@tcp(127.0.0.1:3306)/beegodb?charset=utf8"
	beego.Run()
}

