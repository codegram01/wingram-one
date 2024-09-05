package post

import "database/sql"

func (rs *Resource) DbCreate(postReq *Post) (*Post, error) {
	var post Post

	queryStr := `
		INSERT INTO post 
			(title, description, content, profile_id)
		VALUES 
			($1, $2, $3, $4) 
		RETURNING id, title, description, content, profile_id
	`
	row := rs.Db.QueryRow(
		queryStr,
		postReq.Title,
		postReq.Description,
		postReq.Content,
		postReq.ProfileId,
	)

	err := row.Scan(
		&post.Id,
		&post.Title,
		&post.Description,
		&post.Content,
		&post.ProfileId,
	)

	return &post, err
}

func (rs *Resource) DbUpdate(id int64, postReq *Post) (*Post, error) {
	var post Post

	queryStr := `
		UPDATE post
		SET
			title = $2,
			description = $3,
			content = $4
		WHERE id = $1
		RETURNING id, title, description, content, profile_id
	`
	row := rs.Db.QueryRow(
		queryStr,
		id,
		postReq.Title,
		postReq.Description,
		postReq.Content,
	)

	err := row.Scan(
		&post.Id,
		&post.Title,
		&post.Description,
		&post.Content,
		&post.ProfileId,
	)

	return &post, err
}

func (rs *Resource) DbDetail(id int64) (*Post, error) {
	var post Post

	queryStr := `
		SELECT
			id, title, description, content, profile_id
		FROM post
		WHERE id = $1
	`
	row := rs.Db.QueryRow(queryStr, id)

	err := row.Scan(
		&post.Id,
		&post.Title,
		&post.Description,
		&post.Content,
		&post.ProfileId,
	)

	return &post, err
}

func (rs *Resource) DbDelete(id int64) error {
	queryStr := `DELETE FROM post WHERE id = $1`
	_, err := rs.Db.Exec(queryStr, id)

	return err
}

func (rs *Resource) DbList() ([]*Post, error) {
	var posts []*Post

	query := `
		SELECT
			id, title, description, content, profile_id
		FROM
			post
	`
	err := rs.Db.Query(query, func(rows *sql.Rows) error {
		var post Post

		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.ProfileId,
		)

		if err != nil {
			return err
		}
		posts = append(posts, &post)
		return nil
	})

	return posts, err
}
