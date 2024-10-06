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
	user.Role = &entity.Role{}
	err := u.db.QueryRow("SELECT U.id, U.login, U.nick, U.password, U.created_at, U.activated_at, R.id, R.name "+
		"FROM user_api.user AS U JOIN user_api.role R ON U.role = R.id  "+
		"WHERE nick = $1 LIMIT 1", nickname).
		Scan(&user.Id, &user.Login, &user.Nick, &user.Password, &user.CreatedAt, &user.ActivatedAt, &user.Role.Id, &user.Role.Name)
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
