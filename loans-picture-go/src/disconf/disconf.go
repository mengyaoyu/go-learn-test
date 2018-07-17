package disconf

import (
	p "github.com/magiconair/properties"
	"github.com/gorilla/websocket"
	"net/url"
	"encoding/json"
	"disconf/recws"
	"time"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type Config struct {
	Host    string
	Path    string
	App     string
	Env     string
	Version string
	Key     string
	KeyType string
	success bool

	List interface{}
}

type WatchedChange struct {
	App     string
	Env     string
	Version string
	Key     string
	KeyType string
	Value   string
	Success bool
}

func ConvertToProperties(content string) properties {

	return properties{p.MustLoadString(content)}

}

func GetFile(config Config) string {

	return getDisconfContent(config)
}
func GetItem(config Config) (value, message string, status float64) {

	content := getDisconfContent(config)
	var m map[string]interface{}
	json.Unmarshal([]byte(content), &m)

	return m["value"].(string), m["message"].(string), m["status"].(float64)
}

func getDisconfContent(config Config) string {
	url := "http://" + config.Host +
		"/" + config.Path + "/api/config/" + config.KeyType + "?app=" +
		config.App + "&env=" + config.Env + "&version=" + config.Version + "&key=" + config.Key
	content := getStringFromUrl(url)

	logs.Info("GetDisconfContent,app=%s,env=%s,version=%s,key=%s,content=%s\n",
		config.App, config.Env, config.Version, config.KeyType, content)

	return content
}

//func Watch(content string, config Config, f func(message [] byte)) {
func Watch(content string, config Config, f func(watchConfig WatchedChange)) {
	logs.Info("content", content)
	m := map[string]string{"app": config.App,
		"env": config.Env,
		"version": config.Version,
		"key": config.Key,
		"content": content,
		"type": config.KeyType}
	data, err := json.Marshal(m)
	logs.Info("err:", err)
	logs.Info("data:", string(data))
	if err != nil {
		panic(err)
	}

	u := url.URL{Scheme: "ws", Host: config.Host, Path: config.Path + "/api/changeMonitor"}
	conn := &recws.RecConn{
		RecIntvlMin:      0,
		RecIntvlMax:      0,
		RecIntvlFactor:   0,
		HandshakeTimeout: 0,
		NonVerbose:       false,
		Conn:             &websocket.Conn{},
		InitMessage:      nil,
		InitFunc: func(conn *recws.RecConn) {
			conn.WriteMessage(websocket.TextMessage, data)
		},
	}

	conn.Dial(u.String(), nil)

	for {
		select {
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				time.Sleep(2 * time.Second)
				continue
			}
			logs.Info("disconf配置更改后获取信息: %s\n", message)
			var watchedConfig WatchedChange
			json.Unmarshal(message, &watchedConfig)
			f(watchedConfig)
		}
	}
}
