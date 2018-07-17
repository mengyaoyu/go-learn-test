package main

import (
	_ "routers"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"disconf"
)

func main() {
	//日志处理，一次引入，全局使用，无需重新引入
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLevel(logs.LevelDebug)
	//logs.SetLogger(logs.AdapterFile, `{"filename":"catalina.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//输出log时能显示输出文件名和行号（非必须）
	logs.EnableFuncCallDepth(true)

	//disconf start
	var host = "192.168.30.10:8080"
	config := disconf.Config{Host:host, Path:"", App:"loans-picture-web", Env:"online", Version: "4.0.0", Key:"loans-picture.properties", KeyType: "file"}
	disconf.DisconfDown(config)
	//disconf end

	beego.Run();
}