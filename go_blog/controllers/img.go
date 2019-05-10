package controllers

import (

	"time"
	"github.com/astaxie/beego"
	"path"
)

// 图片处理
// 图片处理
type ImageController struct{
	beego.Controller
}



func (c *ImageController) ShowTest( ) {

	//beego.Info(c.Ctx.Request)

	c.TplName = "test.html"
}







func (c *ImageController) GetTest( ) {

	//c.TplName = "test.html"

	//id,_ := c.GetInt("id")
	//artiName := c.GetString("articleName")
	//content := c.GetString("content")
	//f,h,err:=c.GetFile("uploadname")

	// beego.Info(c.Ctx.Request.Method)  //获取请求方法

	//beego.Info(c.Ctx.Request.Method)  //获取请求方法
	//beego.Info(c.Ctx.Request.AddCookie)  //获取请求方法



	file, fileHeader, err := c.GetFile("upload")
	file.Close()

	if err != nil {
		return
	}
	defer file.Close()


	////1.要限定格式
	fileext := path.Ext(fileHeader.Filename)
	//if fileext != ".jpg" && fileext != "png"{
	//	beego.Info("上传文件格式错误")
	//	return
	//}
	////2.限制大小
	//if h.Size > 50000000 {
	//	beego.Info("上传文件过大")
	//	return
	//}

	//3.需要对文件重命名，防止文件名重复
	filename := time.Now().Format("2006-01-02 15:04:05") + fileext  //6-1-2 3:4:5

	   beego.Info(c.Ctx.Request.URL)
	 beego.Info(c.Ctx.Request.URL +filename)

	if err != nil{
		beego.Info("上传文件失败")
		return
	}else {
		c.SaveToFile("upload","./static/img/"+filename)
	}



	//
	//data := map[string]interface{}{
	//	"uploaded": 1,
	//	"fileName": filename,
	//	"url":      cdnDomain + path,
	//}
	//b, err := json.Marshal(data)
	//if err != nil {
	//	return
	//}
	//


	c.Redirect("/test",302)



}
