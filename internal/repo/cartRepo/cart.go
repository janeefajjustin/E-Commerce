package cartRepo

import "database/sql"

type CartRepo struct {
	db *sql.DB
}

func NewCartRepo(db *sql.DB) *CartRepo{
	return &CartRepo{db:db}
}