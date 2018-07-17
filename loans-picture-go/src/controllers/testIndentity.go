package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"mime/multipart"
	"bytes"
	"io"
	"github.com/astaxie/beego/logs"
)

type TestIdentityController  struct {
	beego.Controller
}

// @router /upload/id/card [post]
func (c *TestIdentityController) UploadIdCardtest1() {

	var client  = http.Client{}
	var url  = "http://127.0.0.1:8080/upload/id/cardtest2"

	f, h, _ := c.GetFile("f1")
	f2, h2, _ := c.GetFile("f2")
	defer f.Close()
	defer f2.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part,_:=writer.CreateFormFile("f1",h.Filename)
	logs.Info(h.Filename)
	io.Copy(part,f)
	part2,_:=writer.CreateFormFile("f2",h2.Filename)
	io.Copy(part2,f2)
	var contentType  = writer.FormDataContentType()
	writer.Close()
	resp, _ := client.Post(url,contentType,body)
	defer resp.Body.Close()
	logs.Info(resp.StatusCode)

	return
}

// @router /upload/id/cardt [post]
func (c *TestIdentityController) UploadIdCardtest2() {

	f, h, _ := c.GetFile("f1")
	logs.Info(h.Filename)
	f2, h2, _ := c.GetFile("f2")
	logs.Info(h2.Size)



	defer f.Close()
	f2.Close()
	//f3 := c.GetString("f3")
	//logs.Info(f3)

	c.ServeJSON()
	c.StopRun()
	return
}
