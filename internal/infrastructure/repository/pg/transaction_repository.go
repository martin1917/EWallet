package pg

import (
	"database/sql"
	"fmt"

	"github.com/martin1917/EWallet/internal/core/transaction"
)

var _ transaction.Repository = &transactionRepository{}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

// Create implements transaction.Repository
func (r *transactionRepository) Create(t *transaction.Transaction) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	reduceQuery := fmt.Sprintf("UPDATE %s SET balance = balance - $1 WHERE id = $2", walletTable)
	_, err = tx.Exec(reduceQuery, t.Amount, t.From)
	if err != nil {
		tx.Rollback()
		return err
	}

	addQuery := fmt.Sprintf("UPDATE %s SET balance = balance + $1 WHERE id = $2", walletTable)
	_, err = tx.Exec(addQuery, t.Amount, t.To)
	if err != nil {
		tx.Rollback()
		return err
	}

	newTransactionQuery := fmt.Sprintf("INSERT INTO %s (id, from_wallet, to_wallet, amount, date_time) VALUES ($1, $2, $3, $4, $5)", transactionTable)
	_, err = tx.Exec(newTransactionQuery, t.Id, t.From, t.To, t.Amount, t.DateTime)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetLast implements transaction.Repository
func (r *transactionRepository) GetLast(count int) ([]transaction.Transaction, error) {
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY date_time DESC LIMIT $1", transactionTable)
	rows, err := r.db.Query(query, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]transaction.Transaction, 0, count)
	for rows.Next() {
		transaction := transaction.Transaction{}
		rows.Scan(&transaction.Id, &transaction.From, &transaction.To, &transaction.Amount, &transaction.DateTime)
		transactions = append(transactions, transaction)
	}

	return transactions[:], nil
}
