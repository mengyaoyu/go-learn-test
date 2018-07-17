package controllers

import (
	"github.com/astaxie/beego"
	"disconf"
)

type MainController  struct {
	beego.Controller
}

func (c *MainController)Get()  {
	conf := new(disconf.Config)
	conf.Host = "127.0.0.1"
	conf.KeyType="file"
	conf.Version="4.0.0"
	conf.App="loan-identity-web"
	conf.Env="online"
	conf.Key = "application.properties"
	c.ServeJSON()
	c.StopRun()
	return
}
