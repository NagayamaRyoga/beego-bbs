package controllers

import (
	"html/template"

	"github.com/NagayamaRyoga/beego-bbs/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// MainController はメインページの制御を行う
type MainController struct {
	beego.Controller
}

// Get は投稿一覧を表示する
// @router	/	[get]
func (c *MainController) Get() {
	// すべての投稿を取得する
	posts, err := models.GetAll(-1, 0)
	if err != nil {
		c.Abort("500")
	}

	// フォームの初期値を設定する
	c.Data["Author"] = ""
	c.Data["Body"] = ""

	c.Data["Posts"] = posts
	c.Data["Error"] = err
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "index.tpl"
}

// Post は投稿を行う
// @Param	author	{string}	string	true	投稿者名
// @Param	body	{string}	string	true	投稿の本文
// @router	/	[post]
func (c *MainController) Post() {
	var post models.Post
	if err := c.ParseForm(&post); err != nil {
		c.Abort("500")
	}

	// リクエストのバリデーションを行う
	validation := validation.Validation{}
	if isValid, err := validation.Valid(&post); !isValid || err != nil {
		c.Abort("400")
	}

	// 投稿を追加する
	if _, err := models.Add(&post); err != nil {
		c.Abort("500")
	}

	// すべての投稿を取得する
	posts, err := models.GetAll(-1, 0)
	if err != nil {
		c.Abort("500")
	}

	// フォームの初期値を設定する
	c.Data["Author"] = post.Author
	c.Data["Body"] = ""

	c.Data["Posts"] = posts
	c.Data["Error"] = err
	c.Data["ValidationErrors"] = validation.ErrorsMap
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "index.tpl"
	post.CreatedAt.Local()
}
