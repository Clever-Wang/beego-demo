package main

import (
	_ "beego-request/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.MaxMemory = 1<<22
	beego.Run()
}

