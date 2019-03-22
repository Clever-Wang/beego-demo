package routers

import (
	"beego-controller/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api", &controllers.BaseAdminRouter{})
	beego.Router("/stop", &controllers.RController{})
}
