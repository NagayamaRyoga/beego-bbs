package controllers

import (
	"github.com/astaxie/beego"
)

// ErrorController はエラー画面の表示を行う
type ErrorController struct {
	beego.Controller
}

// Error400 はバリデーションエラーを表示する
func (c *ErrorController) Error400() {
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.Data["Message"] = "400 Bad Request"
	c.TplName = "error.tpl"
}

// Error422 はリクエストエラーを表示する
func (c *ErrorController) Error422() {
	c.Ctx.ResponseWriter.WriteHeader(422)
	c.Data["Message"] = "422 Unprocessable Entity"
	c.TplName = "error.tpl"
}

// Error500 はサーバーエラーを表示する
func (c *ErrorController) Error500() {
	c.Ctx.ResponseWriter.WriteHeader(500)
	c.Data["Message"] = "500 Internal Server Error"
	c.TplName = "error.tpl"
}
