package services

import sqliteadmin "templtest/internal/services/sqlite-admin"

type DB interface {
	sqliteadmin.DB
}

type RegisterService struct {
	*sqliteadmin.SQLiteAdminService
}

func New(db DB) *RegisterService {
	sqliteAdminS := sqliteadmin.New(db)
	return &RegisterService{SQLiteAdminService: sqliteAdminS}
}
