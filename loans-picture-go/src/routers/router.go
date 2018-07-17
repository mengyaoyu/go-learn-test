package routers

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IdentityController{})
	beego.Router("/", &controllers.MainController{})
}
