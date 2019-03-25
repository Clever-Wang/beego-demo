package main

import (
	_ "beego-restful/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm" //引入 beego orm
	_"github.com/go-sql-driver/mysql" //初始化mysql驱动
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//注册数据库信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beegodb?charset=utf8")
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
	beego.Run()
}
