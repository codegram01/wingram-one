package token

import "github.com/codegram01/wingram-one/database"

type Token struct {
	Token_id   int64
	Account_id int64
}

type Resource struct {
	*database.Db
}
