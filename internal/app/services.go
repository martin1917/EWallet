package app

import (
	"github.com/martin1917/EWallet/internal/app/commands"
	"github.com/martin1917/EWallet/internal/app/queries"
	"github.com/martin1917/EWallet/internal/infrastructure/repository"
)

type Queries struct {
	GetBalanceHandler queries.GetBalanceRequestHandler
	GetLastHandler    queries.GetLastRequestHandler
}

type Commands struct {
	SendModeyHandler commands.SendModeyRequestHandler
}

type Services struct {
	Queries
	Commands
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		Queries: Queries{
			GetBalanceHandler: queries.NewGetBalanceRequestHandler(repositories.WalletRepository),
			GetLastHandler:    queries.NewGetLastRequestHandler(repositories.TransactionRepository),
		},
		Commands: Commands{
			SendModeyHandler: commands.NewSendModeyRequestHandler(repositories.TransactionRepository, repositories.WalletRepository),
		},
	}
}
