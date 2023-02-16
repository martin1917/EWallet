package repository

import (
	"database/sql"

	"github.com/martin1917/EWallet/internal/core/transaction"
	"github.com/martin1917/EWallet/internal/core/wallet"
	"github.com/martin1917/EWallet/internal/infrastructure/repository/pg"
)

type Repositories struct {
	WalletRepository      wallet.Repository
	TransactionRepository transaction.Repository
}

func NewPgRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		WalletRepository:      pg.NewWalletRepository(db),
		TransactionRepository: pg.NewTransactionRepository(db),
	}
}
