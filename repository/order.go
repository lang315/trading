package repository

import (
	"github.com/go-pg/pg"
	"projects/trading/models"
)

func InsertOrder(db *pg.DB, order *models.Order) {
	err := db.Insert(order)
	if err!=nil {
		println("Err insert order: " + err.Error())
	}
}
