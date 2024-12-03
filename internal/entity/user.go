package entity

type User struct {
	Guid   string
	Active bool

	Role Role

	CreatedAt      string
	ActivatedAt    string
	ActivationCode string

	Login    string
	Nickname string
	Email    string
	Password string
	Avatar   string
}