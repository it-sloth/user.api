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

func (u User) GetUser(login string) (*entity.User, error) {
	var user entity.User
	user.Role = &entity.Role{}

	err := u.db.QueryRow(`
		SELECT U.id, 
		       U.login, 
		       U.nick, 
		       U.password, 
		       U.created_at, 
		       U.activated_at, 
		       R.id,
		       R.name 
		FROM user_api.user AS U 
		JOIN user_api.role R ON U.role = R.id 
		WHERE LOWER(U.login) = LOWER($1) LIMIT 1`, login).
		Scan(&user.Id, &user.Login, &user.Nick, &user.Password, &user.CreatedAt, &user.ActivatedAt, &user.Role.Id, &user.Role.Name)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) CheckUser(login string) (string, error) {
	var guid string

	err := u.db.QueryRow("SELECT id FROM user_api.user WHERE lower(login) = lower($1) LIMIT 1", login).Scan(&guid)
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
