package account

import (
	"github.com/codegram01/wingram-one/database"
)

type Account struct {
	Id    int    `json:"id"`
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
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	ProfileId int    `json:"profile_id"`
}

type Auth struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type Resource struct {
	database.Db
}
