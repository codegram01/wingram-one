package account

import (
	"database/sql"
)

func (rs *Resource) DbCreate(accReq *AccountReq) (*Account, error) {
	var acc Account

	queryStr := `
		INSERT INTO account 
			(email, password)
		VALUES 
			($1, $2) 
		RETURNING id, email
	`
	row := rs.Db.QueryRow(queryStr, accReq.Email, accReq.Password)

	err := row.Scan(
		&acc.Id,
		&acc.Email,
	)

	return &acc, err
}

// careful with this function
// it get have password
func (rs *Resource) DbDetailAuth(email string) (*AccountAuth, error) {
	var acc AccountAuth

	queryStr := `
		SELECT
			id, email, password
		FROM
			account
		WHERE 
			email = $1
	`
	row := rs.Db.QueryRow(queryStr, email)

	err := row.Scan(
		&acc.Id,
		&acc.Email,
		&acc.Password,
	)

	return &acc, err
}

func (rs *Resource) DbInfoDetail(id int) (*AccountInfo, error) {
	var acc AccountInfo

	queryStr := `
		SELECT
			a.id, a.email, p.name, p.id
		FROM
			account a
			INNER JOIN profile p ON a.id = p.account_id
		WHERE 
			a.id = $1
	`
	row := rs.Db.QueryRow(queryStr, id)

	err := row.Scan(
		&acc.Id,
		&acc.Email,
		&acc.Name,
		&acc.ProfileId,
	)

	return &acc, err
}

func (rs *Resource) DbInfoList() ([]*AccountInfo, error) {
	var accounts []*AccountInfo

	query := `
		SELECT
			a.id, a.email, p.name, p.id
		FROM
			account a
			INNER JOIN profile p ON a.id = p.account_id
	`
	err := rs.Db.Query(query, func(rows *sql.Rows) error {
		var acc AccountInfo

		err := rows.Scan(
			&acc.Id,
			&acc.Email,
			&acc.Name,
			&acc.ProfileId,
		)

		if err != nil {
			return err
		}
		accounts = append(accounts, &acc)
		return nil
	})

	return accounts, err
}

func (rs *Resource) GetAccountByToken(idToken int) (int, error) {
	var idAccount int
	row := rs.Db.QueryRow(`
		SELECT account_id
		FROM token
		WHERE id = $1`, idToken)

	err := row.Scan(
		&idAccount,
	)

	return idAccount, err
}
