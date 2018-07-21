package repository

import (
	"github.com/go-pg/pg"
	"projects/trading/models"
)

func InsertOrder(db *pg.DB, order *models.Order) {
	_,err := db.Model(&order).Insert()
	if err!=nil {
		println("Err insert order: " + err.Error())
	}
}
