package models

import "github.com/satori/go.uuid"

type Balance struct {
	tableName struct{}  `sql:"balance"`
	ID int
	USD float64
	EUR float64
	UserID uuid.UUID
}
