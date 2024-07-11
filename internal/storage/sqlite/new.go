package sqlite

import (
	"fmt"
	"templtest/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sqlx.DB
}

func New() (*Repository, error) {
	db, err := sqlx.Open("sqlite3", config.MainSQLitePath)
	if err != nil {
		return &Repository{}, err
	}

	err = db.Ping()
	if err != nil {
		return &Repository{}, err
	}

	newDB := &Repository{
		db: db,
	}
	fmt.Println("db connected")
	fmt.Println(config.MainSQLitePath)
	return newDB, nil
}

func (repo *Repository) Close() {
	repo.db.Close()
}
