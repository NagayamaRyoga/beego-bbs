package routers

import (
	"github.com/NagayamaRyoga/beego-bbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.ErrorController(&controllers.ErrorController{})
}
