package sqlite

import (
	"templtest/internal/entities"

	sq "github.com/Masterminds/squirrel"
)

func (repo *Repository) Users() ([]entities.User, error) {
	sql, arg, err := sq.Select("id", "name", "full_name", "email", "avatar", "language", "created_at").From("users").ToSql()
	if err != nil {
		return []entities.User{}, err
	}
	result := make([]entities.User, 0)
	err = repo.db.Select(&result, sql, arg...)
	if err != nil {
		return []entities.User{}, err
	}
	return result, nil
}
