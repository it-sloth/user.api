package entity

type User struct {
	Guid   string
	Active bool

	Role Role

	CreatedAt      string
	UpdatedAt      string
	ActivatedAt    string
	ActivationCode string

	Login    string
	Nickname string
	Email    string
	Password string
	Icon     string
}