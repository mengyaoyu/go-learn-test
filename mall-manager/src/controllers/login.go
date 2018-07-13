package controllers

import (
	"github.com/astaxie/beego"
	"models"
	"commons"
	"github.com/astaxie/beego/logs"
)

type LoginController  struct {
	beego.Controller
}

// @router /login [get]
func (c *LoginController) Login() {
	c.TplName = "signin.html"
}

// @router /user/add/user [get]
func (c *LoginController) AddUser() {

	user := new(models.UserInfo)
	user.UserCode = "admin"
	user.UserPwd = commons.EncoderByMd5("123")
	user.UserName = "管理员"
	user.CreateUser="sys"
	user.UpdateUser="sys"
	cnt ,_ :=models.SaveUserInfo(user)
	logs.Info(cnt)
	c.TplName = "signin.html"

}

// @router /login/sign [get]
func (c *LoginController) Sign() {
	code := c.Input().Get("code")
	user, _ := models.UserGetByUserCode(code)
	if user == nil {
		c.Data["errorMsg"] = "用户名不正确"
		c.TplName = "signin.html"
		return
	}
	if user.UserPwd != commons.EncoderByMd5(c.Input().Get("pwd")) {
		c.Data["errorMsg"] = "密码不正确"
		c.TplName = "signin.html"
		return
	}
	c.SetSession(commons.CURRENT_USER, user)

	redirectUrl := "index.html"
	if c.Input().Get("rurl") != "" {
		redirectUrl = c.Input().Get("rurl")
	}
	c.Ctx.Redirect(302, redirectUrl)
}