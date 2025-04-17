package repositories

import (
	"database/sql"
)

type Entity interface{}

type Repository struct {
	DB *sql.DB
}

func NewRepository (db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Execute(query string, args ...interface{}) error {
	_, err := r.DB.Exec(query, args...)
	return err
}


func (r *Repository) Query(query string, handler func(*sql.Rows) error, args ...interface{}) error {
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return handler(rows)
}

func (r *Repository) QueryRow(query string, args ...interface{}) *sql.Row {
	return r.DB.QueryRow(query, args...)
}