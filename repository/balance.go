package repository

import (
	"github.com/satori/go.uuid"
	"github.com/go-pg/pg"
	"projects/trading/models"
)

func SelectBalance(db *pg.DB, userID uuid.UUID) *models.Balance{
	b := &models.Balance{}
	db.Model(&b).Where("user_id=?", userID).Select()
	return b
}
