package entities

import "time"

type User struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	FullName  string    `db:"full_name" json:"full_name"`
	Email     string    `db:"email" json:"email"`
	Avatar    string    `db:"avatar" json:"avatar"`
	Language  string    `db:"language" json:"language"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type WordBasic struct {
	ID           string `db:"id" json:"id"`
	Title        string `db:"title" json:"title"`
	Description  string `db:"description" json:"description"`
	FromLanguage string `db:"from_language" json:"from_language"`
	ToLanguage   string `db:"to_language" json:"to_language"`
	Type         string `db:"type" json:"type"`
}

type Word struct {
	WordBasic
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    string    `db:"user_id" json:"user_id"`
}
