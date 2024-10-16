// Package repository implements functionality of working with DB
package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/entity"
)

type User struct {
	db *sql.DB
}

// GetUser returns all fields for user from db, you may use other methods to get user with fever fields.
func (u User) GetUser(guid string) (*entity.User, error) {
	var user entity.User
	user.Role = &entity.Role{}

	err := u.db.QueryRow(`
		SELECT U.guid, U.login, U.nickname, U.email, U.active, U.password, U.created_at,
		       U.activated_at, R.id, R.name
		FROM user_api.user AS U JOIN user_api.role R ON U.role = R.id
		WHERE U.guid = $1 LIMIT 1`, guid).
		Scan(&user.Guid, &user.Login, &user.Nickname, &user.Email, &user.Active, &user.Password,
			&user.CreatedAt, &user.ActivatedAt, &user.Role.Id, &user.Role.Name)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CheckUser returns guid if user with specified login and mail exists
func (u User) CheckUser(login string, email string) (string, error) {
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

func NewUserRepository(env *config.Env) *User {
	connection, err := sql.Open("postgres", env.DbDsn)
	if err != nil {
		panic(err.Error())
	}

	if err := connection.Ping(); err != nil {
		panic(err)
	}

	return &User{
		db: connection,
	}
}
