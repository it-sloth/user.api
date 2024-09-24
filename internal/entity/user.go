package entity

import "time"

type User struct {
	Id          string    `db:"id"`
	Login       string    `db:"login"`
	Nick        string    `db:"nick"`
	Password    string    `db:"password"`
	CreatedAt   time.Time `db:"created_at"`
	ActivatedAt time.Time `db:"activated_at"`
}
