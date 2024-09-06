package profile

import (
	"text/template"

	"github.com/codegram01/wingram-one/database"
)

type Profile struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	AccountId int64  `json:"account_id"`
}

type ProfileReq struct {
	Name      string `json:"name"`
	AccountId int64  `json:"account_id"`
}

type Resource struct {
	*database.Db
	*template.Template
}
