package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers:LoginController"] = append(beego.GlobalControllerRouter["controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:LoginController"] = append(beego.GlobalControllerRouter["controllers:LoginController"],
		beego.ControllerComments{
			Method: "Sign",
			Router: `/login/sign`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:LoginController"] = append(beego.GlobalControllerRouter["controllers:LoginController"],
		beego.ControllerComments{
			Method: "AddUser",
			Router: `/user/add/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
