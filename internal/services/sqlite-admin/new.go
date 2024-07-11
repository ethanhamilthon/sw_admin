package sqliteadmin

import "templtest/internal/entities"

type DB interface {
	Users() ([]entities.User, error)
}

type SQLiteAdminService struct {
	db DB
}

func New(db DB) *SQLiteAdminService {
	return &SQLiteAdminService{db}
}

func (s *SQLiteAdminService) AllUsers() ([]entities.User, error) {
	return s.db.Users()
}
