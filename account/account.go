package account

import (
	"github.com/codegram01/wingram-one/database"
	"github.com/codegram01/wingram-one/template"
)

type Account struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type AccountAuth struct {
	Account
	Password string `json:"_"`
}

type AccountReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountInfo struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	ProfileId int64  `json:"profile_id"`
}

type Resource struct {
	*database.Db
	*template.Template
}
