package entity

type User struct {
	UserId      string `json:"user_id" `
	UserName    string `json:"username" `
	Email       string `json:"email" `
	PhoneNumber string `json:"phone_number"`
	Hash        string `json:"_"`
}
