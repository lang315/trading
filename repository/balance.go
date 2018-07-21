package repository

import (
	"github.com/satori/go.uuid"
	"github.com/go-pg/pg"
	"projects/trading/models"
)

func SelectBalance(db *pg.DB, userID uuid.UUID) *models.Balance {
	b := &models.Balance{}
	db.Model(&b).Where("user_id=?", userID).Select()
	return b
}

func UpdateBalance(db *pg.DB, balance *models.Balance) {
	_, err := db.Model(&balance).Set("USD=?", balance.USD).Set("EUR=?", balance.EUR).
		Where("user_id=?", balance.UserID).Update()
	if err != nil {
		println("Error")
		println(err.Error())
	}
}
