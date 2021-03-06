package repository

import (
	"github.com/go-pg/pg"
	"projects/trading/models"
	"github.com/satori/go.uuid"
)

type UserRepository struct {
	User *models.User
}

func (self *UserRepository) IsAlreadyAccount(db *pg.DB) (check bool, user *models.User) {
	var u models.User
	err := db.Model(&u).Where("email=?", self.User.Email).Select()
	//println(u.Email)
	if err != nil {
		return false, &u
	}
	return true, &u
}

func (self *UserRepository) SignUpAccount(db *pg.DB) {
	id, _ := uuid.NewV4()
	self.User.ID = id
	u := models.User{
		ID:       id,
		Email:    self.User.Email,
		Password: self.User.Password,
		Fullname: self.User.Fullname,
		Type:     1,
	}
	err := db.Insert(&u)
	if err != nil {
		println(err.Error())
		return
	}
	println("OK")
}
