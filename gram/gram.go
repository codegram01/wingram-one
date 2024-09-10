package gram

import (
	"github.com/codegram01/wingram-one/database"
	"github.com/codegram01/wingram-one/template"
)

type Gram struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	ParentId   int64  `json:"parent_id"`
	AccountId   int64  `json:"account_id"`
}

type Resource struct {
	*database.Db
	*template.Template
}
