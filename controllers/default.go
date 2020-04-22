package controllers

import (
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

	// フォームの初期値を設定する
	c.Data["Author"] = ""
	c.Data["Body"] = ""

	c.Data["Posts"] = posts
	c.Data["Error"] = err
	c.TplName = "index.tpl"
}

// Post は投稿を行う
// @Param	author	{string}	string	true	投稿者名
// @Param	body	{string}	string	true	投稿の本文
// @router	/	[post]
func (c *MainController) Post() {
	var post models.Post
	_ = c.ParseForm(&post)

	// リクエストのバリデーションを行う
	validation := validation.Validation{}
	isValid, _ := validation.Valid(&post)

	// 投稿を追加する
	if isValid {
		_, _ = models.Add(&post)
	}

	// すべての投稿を取得する
	posts, err := models.GetAll(-1, 0)

	// フォームの初期値を設定する
	if isValid {
		c.Data["Author"] = post.Author
		c.Data["Body"] = ""
	} else {
		c.Data["Author"] = post.Author
		c.Data["Body"] = post.Body
	}

	c.Data["Posts"] = posts
	c.Data["Error"] = err
	c.Data["ValidationErrors"] = validation.ErrorsMap
	c.TplName = "index.tpl"
}
