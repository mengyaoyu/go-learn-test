package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers:IdentityController"] = append(beego.GlobalControllerRouter["controllers:IdentityController"],
		beego.ControllerComments{
			Method: "UploadIdCard",
			Router: `/upload/id/card`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:IdentityController"] = append(beego.GlobalControllerRouter["controllers:IdentityController"],
		beego.ControllerComments{
			Method: "UploadIdCardt",
			Router: `/upload/id/cardt`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
