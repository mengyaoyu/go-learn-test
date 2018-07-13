package main

import (
	_ "routers"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"models"
	"fmt"
)

func main() {
	//日志处理，一次引入，全局使用，无需重新引入
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLevel(logs.LevelDebug)
	//logs.SetLogger(logs.AdapterFile, `{"filename":"catalina.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//输出log时能显示输出文件名和行号（非必须）
	logs.EnableFuncCallDepth(true)
	beego.Run();
}
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbDataBase := beego.AppConfig.String("db.database")
	dbUserName := beego.AppConfig.String("db.username")
	dbPwd := beego.AppConfig.String("db.pwd")

	conn :=  dbUserName + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDataBase + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", conn, 30, 30)//注册默认数据库
	orm.RegisterModel(new(models.UserInfo))
	orm.RunSyncdb("default", false, true)
	fmt.Println("database init is complete . \n Please restart the application")
}