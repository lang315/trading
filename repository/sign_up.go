package repository

import (
	"github.com/go-pg/pg"
	"projects/trading/models"
	"github.com/satori/go.uuid"
)

type UserRepository struct {
}

func (self *UserRepository) IsAlreadyAccount(db *pg.DB, user *models.User) bool {
	var u models.User
	err := db.Model(&u).Where("email=?", user.Email).Select()
	//println(u.Email)
	if err != nil {
		return false
	}
	return true
}

func (self *UserRepository) SignUpAccount(db *pg.DB, user *models.User) {
	id, _ := uuid.NewV4()
	user.ID = id
	u := models.User{
		ID:       id,
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
	}
	err := db.Insert(&u)
	if err != nil {
		println(err.Error())
		return
	}
	println("OK")
}
