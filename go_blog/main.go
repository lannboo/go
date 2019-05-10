package main

import (
	_ "go_blog/routers"
	_ "go_blog/models"
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	beego.AddFuncMap("showPrePage",HandPrePage)
	beego.AddFuncMap("showNextPage",HandNextPage)
	beego.Run()
}

func HandPrePage(data int)(string)  {
	pageIndex := data -1
	pageIndex1 := strconv.Itoa(pageIndex)

	return pageIndex1

}

func HandNextPage(data int)(int)  {
	pageIndex := data + 1

	return pageIndex

}

