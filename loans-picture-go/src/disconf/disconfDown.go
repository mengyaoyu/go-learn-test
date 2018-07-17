package disconf

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"reflect"
	"github.com/astaxie/beego/config"
)

//
func DisconfDown(config Config) {
	content := GetFile(config)
	p := ConvertToProperties(content)
	logs.Info("p", p)
	if (p.Len() > 0) {
		contentMap := p.Map()
		for k, v := range contentMap {
			beego.AppConfig.Set(k, v)
			logs.Info("k=%v, v=%v\n", k, beego.AppConfig.String(k))
		}

	}
	appConfig := beego.AppConfig
	for _, i := range []interface{}{beego.BConfig, &beego.BConfig.Listen, &beego.BConfig.WebConfig, &beego.BConfig.Log, &beego.BConfig.WebConfig.Session} {
		AssignSingleConfigNew(i, appConfig)
	}

	go Watch(content, config, func(watchConfig WatchedChange) {
		logs.Info("disconf配置更改后新内容：", ConvertToProperties(watchConfig.Value))
	})

}

func AssignSingleConfigNew(p interface{}, ac config.Configer) {
	pt := reflect.TypeOf(p)
	if pt.Kind() != reflect.Ptr {
		return
	}
	pt = pt.Elem()
	if pt.Kind() != reflect.Struct {
		return
	}
	pv := reflect.ValueOf(p).Elem()

	for i := 0; i < pt.NumField(); i++ {
		pf := pv.Field(i)
		if !pf.CanSet() {
			continue
		}
		name := pt.Field(i).Name
		switch pf.Kind() {
		case reflect.String:
			pf.SetString(ac.DefaultString(name, pf.String()))
		case reflect.Int, reflect.Int64:
			pf.SetInt(ac.DefaultInt64(name, pf.Int()))
		case reflect.Bool:
			pf.SetBool(ac.DefaultBool(name, pf.Bool()))
		case reflect.Struct:
		default:
		//do nothing here
		}
	}
}