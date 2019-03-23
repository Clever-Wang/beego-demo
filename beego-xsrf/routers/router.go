package routers

import (
	"beego-xsrf/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/xsrf",&controllers.XSRFController{})
    beego.AutoRouter(&controllers.XSRFJqueryController{})
}
