package queries

import "github.com/martin1917/EWallet/internal/core/transaction"

type GetLastRequest struct {
	Count int
}

type GetLastResponse struct {
	Transactions []transaction.Transaction `json:"transactions"`
}

type GetLastRequestHandler struct {
	repository transaction.Repository
}

func (h GetLastRequestHandler) Handle(req GetLastRequest) (*GetLastResponse, error) {
	res, err := h.repository.GetLast(req.Count)
	return &GetLastResponse{
		Transactions: res,
	}, err
}

func NewGetLastRequestHandler(repository transaction.Repository) GetLastRequestHandler {
	return GetLastRequestHandler{
		repository: repository,
	}
}
