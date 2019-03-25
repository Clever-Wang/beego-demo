package routers

import (
	"beego-errorhandler/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.AutoRouter( &controllers.ErrorHandlerController{})
}
