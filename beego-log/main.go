package main

import (
	_ "beego-log/routers"
	"github.com/astaxie/beego"
)

func main() {

	//日志配置
	//设置文件输出的文件路径
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	//删除控制台输出日志,默认输出到控制台
	beego.BeeLogger.DelLogger("console")
	//设置日志级别
	beego.SetLevel(beego.LevelInformational)
	//输出日志行号,默认是false
	beego.SetLogFuncCall(true)

	//各种级别的日志打印
	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")

	//启动服务
	beego.Run()
}

