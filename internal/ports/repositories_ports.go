package ports

import "github.com/fedtorres/file-processor/internal/core/domain"

type TransactionsPort interface {
	CreateTransactions(transactions []domain.Transaction) error
	GetAllTransactions() ([]domain.Transaction, error)
}
