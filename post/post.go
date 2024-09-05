package post

import "github.com/codegram01/wingram-one/database"

type Post struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	ProfileId   int64  `json:"profile_id"`
}

type Resource struct {
	database.Db
}
