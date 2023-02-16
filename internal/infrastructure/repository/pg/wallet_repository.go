package pg

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/martin1917/EWallet/internal/core/wallet"
)

var _ wallet.Repository = &walletRepository{}

type walletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) walletRepository {
	return walletRepository{
		db: db,
	}
}

// Create implements wallet.Repository
func (r walletRepository) Create(w *wallet.Wallet) error {
	query := fmt.Sprintf("INSERT INTO %s (id, balance) VALUES ($1, $2)", walletTable)
	_, err := r.db.Exec(query, w.Id, w.Balance)
	return err
}

// GetById implements wallet.Repository
func (r walletRepository) GetById(id uuid.UUID) (wallet.Wallet, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", walletTable)
	row := r.db.QueryRow(query, id)

	var wallet wallet.Wallet
	err := row.Scan(&wallet.Id, &wallet.Balance)
	return wallet, err
}
