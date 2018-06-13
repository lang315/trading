package repository

import (
	"github.com/go-pg/pg"
	"projects/trading/models"
	"github.com/satori/go.uuid"
)

func SelectUser(db *pg.DB, userID uuid.UUID) *models.User {
	var u = &models.User{ID: userID}
	db.Select(u)
	return u
}

