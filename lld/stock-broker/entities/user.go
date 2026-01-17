package entities

type User struct {
	Id   int
	Name string
}

type Address struct {
	Building string
	City     string
}

type AdminUser struct {
	User
}
