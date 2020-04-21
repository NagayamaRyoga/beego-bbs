package main

import (
	_ "github.com/NagayamaRyoga/beego-bbs/models"
	_ "github.com/NagayamaRyoga/beego-bbs/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
