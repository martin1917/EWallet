package errors

import (
	"github.com/google/uuid"
)

type AppError struct {
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

type NotEnoughMoneyError struct {
	Message        string    `json:"message"`
	Id             uuid.UUID `json:"id"`
	CurrentBalance float64   `json:"current_balance"`
	MoneyTransfer  float64   `json:"money_transfer"`
}

func (e *NotEnoughMoneyError) Error() string {
	return "Not enough money"
}
