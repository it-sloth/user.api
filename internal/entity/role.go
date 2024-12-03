package entity

const (
	AdminRole string = "Admin"
	UserRole  string = "User"
)

type Role struct {
	Id   int
	Name string
}