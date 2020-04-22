package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Post struct {
	ID        int64     `orm:"auto;column(id)"`
	Author    string    `orm:"size(128)"                    form:"author"   valid:"Required;MinSize(1);MaxSize(128)"`
	Body      string    `orm:"type(longtext)"               form:"body"     valid:"Required;MinSize(1);MaxSize(1024)"`
	CreatedAt time.Time `orm:"type(datetime);auto_now_add"`
	UpdatedAt time.Time `orm:"type(datetime);auto_now"`
}

func init() {
	orm.RegisterModel(new(Post))
}

// GetAll はすべての投稿を取得する
// limit は取得する投稿数の上限, -1 ならば制限なくすべての投稿を取得する
// offset は取得する投稿の位置
func GetAll(limit int64, offset int64) (posts []Post, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Post)).Limit(limit, offset).OrderBy("-id").All(&posts)
	return
}

// Add はデータベースに投稿を追加する
func Add(post *Post) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(post)
}
