package routers

import (
	"beego-build-url/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/api/list", &controllers.TestController{}, "*:List")
	beego.Router("/person/:last/:first", &controllers.TestController{})
	beego.Router("/params/:first/:second/:third", &controllers.TestController{},"*:Params")
	beego.AutoRouter(&controllers.TestController{})
}
