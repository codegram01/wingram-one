package gram

import "database/sql"

func (rs *Resource) DbCreate(gramReq *Gram) (*Gram, error) {
	var gram Gram

	queryStr := `
		INSERT INTO gram 
			(title, description, content)
		VALUES 
			($1, $2, $3) 
		RETURNING id, title, description, content
	`
	row := rs.Db.QueryRow(
		queryStr,
		gramReq.Title,
		gramReq.Description,
		gramReq.Content,
		// gramReq.ParentId,
		// gramReq.AccountId,
	)

	err := row.Scan(
		&gram.Id,
		&gram.Title,
		&gram.Description,
		&gram.Content,
		// &gram.ParentId,
		// &gram.AccountId,
	)

	return &gram, err
}

func (rs *Resource) DbUpdate(id int64, gramReq *Gram) (*Gram, error) {
	var gram Gram

	queryStr := `
		UPDATE gram
		SET
			title = $2,
			description = $3,
			content = $4
		WHERE id = $1
		RETURNING id, title, description, content
	`
	row := rs.Db.QueryRow(
		queryStr,
		id,
		gramReq.Title,
		gramReq.Description,
		gramReq.Content,
	)

	err := row.Scan(
		&gram.Id,
		&gram.Title,
		&gram.Description,
		&gram.Content,
	)

	return &gram, err
}

func (rs *Resource) DbDetail(id int64) (*Gram, error) {
	var gram Gram

	queryStr := `
		SELECT
			id, title, description, content
		FROM gram
		WHERE id = $1
	`
	row := rs.Db.QueryRow(queryStr, id)

	err := row.Scan(
		&gram.Id,
		&gram.Title,
		&gram.Description,
		&gram.Content,
		// &gram.ParentId,
		// &gram.AccountId,
	)

	return &gram, err
}

func (rs *Resource) DbDelete(id int64) error {
	queryStr := `DELETE FROM gram WHERE id = $1`
	_, err := rs.Db.Exec(queryStr, id)

	return err
}

func (rs *Resource) DbList() ([]*Gram, error) {
	var grams []*Gram

	query := `
		SELECT
			id, title, description, content
		FROM
			gram
	`
	err := rs.Db.Query(query, func(rows *sql.Rows) error {
		var gram Gram

		err := rows.Scan(
			&gram.Id,
			&gram.Title,
			&gram.Description,
			&gram.Content,
			// &gram.ParentId,
			// &gram.AccountId,
		)

		if err != nil {
			return err
		}
		grams = append(grams, &gram)
		return nil
	})

	return grams, err
}
