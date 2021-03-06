package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Order struct {
	tableName struct{}  `sql:"orders"`
	ID        uuid.UUID `sql:",type:uuid,default:uuid_generate_v4()"`
	Symbol string
	Price float64
	OriginQuantity float64
	ExecutedQuantity float64
	Type int
	Time time.Time
	IsWorking bool
	UserID uuid.UUID
}
