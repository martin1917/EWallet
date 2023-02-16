package wallet

import "github.com/google/uuid"

type Wallet struct {
	Id      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}
