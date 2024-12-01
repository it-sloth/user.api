package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func (u *UserRepository) Check(login string, email string) (string, error) {
	var guid string

	err := u.db.QueryRow(`
		SELECT guid FROM user_api.user
		WHERE login = $1 AND email = $2
		LIMIT 1`, login, email).Scan(&guid)

	if err != nil {
		return "", err
	}

	return guid, nil
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}