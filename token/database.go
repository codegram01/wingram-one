package token

import (
	"database/sql"
	"fmt"
)

func (rs *Resource) GetAccountByToken(idToken int64) (int64, error) {
	var idAccount int64
	row := rs.Db.QueryRow(`
		SELECT account_id
		FROM token
		WHERE id = $1`, idToken)

	err := row.Scan(
		&idAccount,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return idAccount, fmt.Errorf("not found token")
		}

		return idAccount, err
	}

	return idAccount, nil
}

func (rs *Resource) GetInfoRefreshToken(idTokenRefresh int) (Token, error) {
	var tokenRefreshInfo Token

	row := rs.Db.QueryRow(`
	SELECT 
		token_id, 
		account_id
	FROM refresh_token
	WHERE id = $1`, idTokenRefresh)

	err := row.Scan(
		&tokenRefreshInfo.Token_id,
		&tokenRefreshInfo.Account_id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return tokenRefreshInfo, fmt.Errorf("not found token")
		}
		return tokenRefreshInfo, err
	}

	return tokenRefreshInfo, nil
}

func (rs *Resource) CreateToken(accountId int64) (int64, error) {
	var tokenId int64

	row := rs.Db.QueryRow(`
		INSERT INTO token (account_id)
		VALUES ($1)
		RETURNING id
	`, accountId)

	err := row.Scan(&tokenId)

	if err != nil {
		return tokenId, err
	}

	return tokenId, err
}

func (rs *Resource) CreateRefreshToken(t Token) (int64, error) {
	var tokenId int64

	row := rs.Db.QueryRow(`
		INSERT INTO refresh_token (account_id, token_id)
		VALUES ($1, $2)
		RETURNING id
	`, t.Account_id, t.Token_id)

	err := row.Scan(&tokenId)

	if err != nil {
		return tokenId, err
	}

	return tokenId, err
}
