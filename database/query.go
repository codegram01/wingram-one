package database

import "database/sql"

func (db *Db) QueryRow(query string, args ...any) *sql.Row {
	return db.Con.QueryRow(query, args...)
}

// RunQuery executes query, then calls f on each row. It stops when there are no
// more rows or f returns a non-nil error.
func (db *Db) Query(query string, f func(*sql.Rows) error, params ...any) error {
	rows, err := db.Con.Query(query, params...)
	if err != nil {
		return err
	}
	_, err = processRows(rows, f)
	return err
}

func processRows(rows *sql.Rows, f func(*sql.Rows) error) (int, error) {
	defer rows.Close()
	n := 0
	for rows.Next() {
		n++
		if err := f(rows); err != nil {
			return n, err
		}
	}
	return n, rows.Err()
}

func (db *Db) Exec(query string, args ...any) (sql.Result, error) {
	return db.Con.Exec(query, args...)
}
