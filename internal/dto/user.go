package dto

import "time"

type User struct {
	Nickname    string    `json:"name"`
	MemberSince time.Time `json:"member_since"`
}

type UserResponse struct {
	Users   []User `json:"users"`
	Version string `json:"version"`
}
