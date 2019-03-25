package routers

import (
	"beego-formvalidation/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/valid", &controllers.FormValidationController{})
    beego.Router("/validStructTag", &controllers.FormValidationController{},"*:ValidStructTag")
}
