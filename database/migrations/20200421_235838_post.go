package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Post_20200421_235838 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Post_20200421_235838{}
	m.Created = "20200421_235838"

	migration.Register("Post_20200421_235838", m)
}

// Run the migrations
func (m *Post_20200421_235838) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE post(`id` int(11) NOT NULL AUTO_INCREMENT,`author` varchar(128) NOT NULL,`body` longtext  NOT NULL,`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Post_20200421_235838) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `post`")
}
