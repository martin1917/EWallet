package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	transactionTable = "transactions"
	walletTable      = "wallets"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
