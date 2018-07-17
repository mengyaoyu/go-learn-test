package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"mime/multipart"
	"bytes"
	"io"
	"github.com/astaxie/beego/logs"
)

type IdentityController  struct {
	beego.Controller
}

// @router /upload/id/card [post]
func (c *IdentityController) UploadIdCard() {
	//获取uid
	uid,_ := c.GetInt8("uid")
	//获取图片a
	af,ah,_ := c.GetFile("a")//af 文件  ah 文件名字
	//获取图片b
	bf,bh,_ := c.GetFile("b")//bf 文件  bh 文件名字

	//http客户端
	var client  = http.Client{}
	//http请求地址
	var url  = "http://127.0.0.1:8080/upload/id/cardt"

	//关闭流
	defer af.Close()
	//关闭流
	defer bf.Close()

	//请求体
	body := &bytes.Buffer{}
	//附件流
	writer := multipart.NewWriter(body)
	//附件a  ah.Filename 不应该为中文负责需要转义
	part,_:=writer.CreateFormFile("a",ah.Filename)
	io.Copy(part,af)
	//附件b  ah.Filename 不应该为中文负责需要转义
	part2,_:=writer.CreateFormFile("idCardReverse",bh.Filename)
	io.Copy(part2,bh)

	//参数
	writer.WriteField("uid",uid)
	//请求头
	var contentType  = writer.FormDataContentType()
	//关闭流
	writer.Close()
	//post请求
	resp, _ := client.Post(url,contentType,body)
	defer resp.Body.Close()
	logs.Info(resp.StatusCode)

	return
}

// @router /upload/id/cardt [post]
func (c *IdentityController) UploadIdCardt() {
	f, h, _ := c.GetFile("a")
	logs.Info(h.Filename)
	f2, h2, _ := c.GetFile("a")
	logs.Info(h2.Filename)

	defer f.Close()
	f2.Close()
	return
}
