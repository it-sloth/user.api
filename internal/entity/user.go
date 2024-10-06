package entity

import "time"

type User struct {
	Id          string
	Login       string
	Nick        string
	Password    string
	Role        *Role
	CreatedAt   time.Time
	ActivatedAt time.Time
}

type Role struct {
	Id   int
	Name string
}
