package models

type User struct {
	ID       string
	Email    string `form:"email"`
	Password string `form:"password"`
	Fullname string `form:"fullname"`
}
