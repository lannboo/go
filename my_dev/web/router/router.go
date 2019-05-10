package router

import (
	"my_dev/web/controller/IndexController"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
