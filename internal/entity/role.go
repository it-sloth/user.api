package entity

const (
	AdminRole string = "admin"
	UserRole  string = "user"
)

type Role struct {
	Id   int
	Name string
}