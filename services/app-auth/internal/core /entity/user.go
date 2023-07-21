package entity

type User struct {
	Name     string
	Email    string
	Id       string
	Phone    string
	Password string
}

type Otp struct {
	PhoneNumber string
	Code        string
}
