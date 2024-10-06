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

func (u User) GetUser(nickname string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow("SELECT id, login, nick, password, created_at, activated_at FROM user_api.users WHERE nick = $1 LIMIT 1", nickname).
		Scan(&user.Id, &user.Login, &user.Nick, &user.Password, &user.CreatedAt, &user.ActivatedAt)
	if err != nil {
		return entity.User{}, err
	}

	return user, err
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
