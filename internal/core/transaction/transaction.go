package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id       uuid.UUID `json:"id"`
	From     uuid.UUID `json:"from"`
	To       uuid.UUID `json:"to"`
	Amount   float64   `json:"amount"`
	DateTime time.Time `json:"date_time"`
}
