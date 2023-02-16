package commands

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/martin1917/EWallet/internal/app/errors"
	"github.com/martin1917/EWallet/internal/core/transaction"
	"github.com/martin1917/EWallet/internal/core/wallet"
)

type SendModeyRequest struct {
	From   uuid.UUID `json:"from"`
	To     uuid.UUID `json:"to"`
	Amount float64   `json:"amount"`
}

type SendModeyRequestHandler struct {
	transaction_repository transaction.Repository
	wallet_repository      wallet.Repository
}

func (h SendModeyRequestHandler) Handle(req SendModeyRequest) error {

	w, err := h.wallet_repository.GetById(req.From)
	if err != nil {
		return &errors.AppError{Message: fmt.Sprintf("Wallet(%s) does not exist", req.From)}
	}

	if req.Amount <= 0.0 {
		return &errors.AppError{Message: "Money transfer must be positive value"}
	}

	if w.Balance < req.Amount {
		return &errors.NotEnoughMoneyError{
			Message:        "Not enough money",
			Id:             w.Id,
			CurrentBalance: w.Balance,
			MoneyTransfer:  req.Amount,
		}
	}

	return h.transaction_repository.Create(
		&transaction.Transaction{
			Id:       uuid.New(),
			From:     req.From,
			To:       req.To,
			Amount:   req.Amount,
			DateTime: time.Now(),
		})
}

func NewSendModeyRequestHandler(
	transaction_repository transaction.Repository,
	wallet_repository wallet.Repository,
) SendModeyRequestHandler {
	return SendModeyRequestHandler{
		transaction_repository: transaction_repository,
		wallet_repository:      wallet_repository,
	}
}
