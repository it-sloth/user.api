package repository

import (
	"database/sql"
	"it-sloth/user.api/internal/entity"

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

func (u *UserRepository) Create(user entity.User) (string, error) {
	var guid string

	err := u.db.QueryRow(`
		INSERT INTO user (login, nickname, email, password, created_at, updated_at, active, role_id) 
           VALUES($1, $2, $3, $4, $5, $6, $7, (SELECT id FROM role WHERE name = $8)) RETURNING guid`,
		   user.Login, user.Nickname, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.Active, user.Role.Name).
		   Scan(&guid)

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