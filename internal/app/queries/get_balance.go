package queries

import (
	"github.com/google/uuid"
	"github.com/martin1917/EWallet/internal/core/wallet"
)

type GetBalanceRequest struct {
	WalletAddress uuid.UUID
}

type GetBalanceRequestHandler struct {
	repository wallet.Repository
}

func (h GetBalanceRequestHandler) Handle(req GetBalanceRequest) (*wallet.Wallet, error) {
	walletModel, err := h.repository.GetById(req.WalletAddress)
	if err != nil {
		return nil, err
	}

	return &wallet.Wallet{
		Id:      walletModel.Id,
		Balance: walletModel.Balance,
	}, nil
}

func NewGetBalanceRequestHandler(repository wallet.Repository) GetBalanceRequestHandler {
	return GetBalanceRequestHandler{
		repository: repository,
	}
}
