package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type TradingHistory struct {
	tableName struct{}  `sql:"trading_history"`
	ID        uuid.UUID `sql:",type:uuid,default:uuid_generate_v4()"`
	Date      time.Location
	Type      int
	Amount    float64
	Price     float64
	Symbol    string
	UserID    uuid.UUID
}
