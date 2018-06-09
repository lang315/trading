package repository

import (
	"github.com/go-pg/pg"
	"projects/trading/models"
)

func IsAlreadyAccount(db *pg.DB, user models.User) bool {
	var u models.User
	db.Model(models.User{}).Where("email=?", user.Email).Select(&u)
	if u.ID!="" {
			return true
	}
	return false
}