package adapters

import (
	"database/sql"

	"github.com/fedtorres/file-processor/internal/core/domain"
)

const (
	createTableQuery = `CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		date DATE,
		amount FLOAT,
		is_credit BOOLEAN,
		is_debit BOOLEAN
	);`

	insertQuery = `INSERT INTO transactions (date, amount, is_credit, is_debit) VALUES ($1, $2, $3, $4);`

	selectQuery = `SELECT id, date, amount, is_credit, is_debit FROM transactions;`
)

type transactionsRepository struct {
	db *sql.DB
}

func NewTransactionsRepository(db *sql.DB) *transactionsRepository {
	return &transactionsRepository{db: db}
}

func (tr transactionsRepository) CreateTransactions(transactions []domain.Transaction) error {
	_, errCreate := tr.db.Exec(createTableQuery)
	if errCreate != nil {
		return domain.ErrExecDb
	}

	// Insert the transactions into the database
	for _, transaction := range transactions {
		_, errInsert := tr.db.Exec(insertQuery, transaction.Date, transaction.Amount, transaction.IsCredit, transaction.IsDebit)
		if errInsert != nil {
			return domain.ErrExecDb
		}
	}

	return nil
}

func (tr transactionsRepository) GetAllTransactions() ([]domain.Transaction, error) {
	rows, errSelect := tr.db.Query(selectQuery)
	if errSelect != nil {
		return nil, domain.ErrQueryDb
	}

	defer rows.Close()

	var transactions []domain.Transaction

	for rows.Next() {
		var transaction domain.Transaction

		errScan := rows.Scan(&transaction.ID, &transaction.Date, &transaction.Amount, &transaction.IsCredit, &transaction.IsDebit)
		if errScan != nil {
			return nil, domain.ErrScanRows
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
