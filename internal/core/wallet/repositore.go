package wallet

import "github.com/google/uuid"

type Repository interface {
	Create(w *Wallet) error
	GetById(id uuid.UUID) (Wallet, error)
}
