package entities

type User struct {
	Id      int
	Name    string
	Phoneno string
}

type Address struct {
	Building string
	City     string
}

type AdminUser struct {
	User
}
