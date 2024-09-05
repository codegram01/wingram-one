package profile

import "github.com/codegram01/wingram-one/database"

type Profile struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	AccountId int    `json:"account_id"`
}

type ProfileReq struct {
	Name      string `json:"name"`
	AccountId int    `json:"account_id"`
}

type Resource struct {
	database.Db
}
