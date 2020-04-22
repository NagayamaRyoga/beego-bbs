package controllers

import (
	"github.com/NagayamaRyoga/beego-bbs/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Posts"], c.Data["Error"] = models.GetAllPost(nil, nil, []string{"id"}, []string{"desc"}, 0, 0)
	c.TplName = "index.tpl"
}
