package entity

import "time"

type User struct {
	Guid        string
	Login       string
	Nickname    string
	Password    string
	Email       string
	Role        *Role
	Active      bool
	CreatedAt   time.Time
	ActivatedAt time.Time
}

type Role struct {
	Id   int
	Name string
}
