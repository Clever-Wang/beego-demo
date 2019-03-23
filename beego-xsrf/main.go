package main

import (
	_ "beego-xsrf/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.BConfig.WebConfig.XSRFExpire = 5  //过期时间，默认1小时

	beego.Run()
}

