package routers

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.LoginController{})
	beego.Router("/", &controllers.MainController{})
}
