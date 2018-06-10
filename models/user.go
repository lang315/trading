package models

import "github.com/satori/go.uuid"

type User struct {
	tableName struct{}  `sql:"account"`
	ID        uuid.UUID `sql:",type:uuid"`
	Email     string    `form:"email"`
	Password  string    `form:"password"`
	Fullname  string    `form:"fullname"`
}
