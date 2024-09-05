package profile

import ()

func (rs *Resource) DbCreate(profileReq *ProfileReq) (*Profile, error) {
	var profile Profile

	queryStr := `
		INSERT INTO profile
			(name, account_id)
		VALUES 
			($1, $2) 
		RETURNING id, name, account_id
	`
	row := rs.Db.QueryRow(queryStr, profileReq.Name, profileReq.AccountId)

	err := row.Scan(
		&profile.Id,
		&profile.Name,
		&profile.AccountId,
	)

	return &profile, err
}
