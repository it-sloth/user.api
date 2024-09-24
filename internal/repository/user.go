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

func (u User) GetUser(login string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT id, login, nick, password, created_at, activated_at FROM user_api.users WHERE login = $1 LIMIT 1", login).
		Scan(&user.Id, &user.Login, &user.Nick, &user.Password, &user.CreatedAt, &user.ActivatedAt)
	if err != nil {
		return entity.User{}, err
	}

	return user, err
}

func NewUserRepository(env *config.Env) *User {
	connection, err := sql.Open("postgres", env.DbDsn)
	if err != nil {
		return nil
	}
	return &User{
		db: connection,
	}
}
