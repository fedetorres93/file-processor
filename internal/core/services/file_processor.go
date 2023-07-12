package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fedtorres/file-processor/internal/core/domain"
	"github.com/fedtorres/file-processor/internal/ports"
)

type fileProcessor struct {
	fileParser            ports.FileParserPort
	summaryGenerator      ports.SummaryGeneratorPort
	emailSender           ports.EmailSenderPort
	transacionsRepository ports.TransactionsPort
}

func NewFileProcessor(
	fileParser ports.FileParserPort,
	summaryGenerator ports.SummaryGeneratorPort,
	emailSender ports.EmailSenderPort,
	transactionsRepository ports.TransactionsPort,
) *fileProcessor {
	return &fileProcessor{
		fileParser:            fileParser,
		summaryGenerator:      summaryGenerator,
		emailSender:           emailSender,
		transacionsRepository: transactionsRepository,
	}
}

func (fp fileProcessor) Process(filePath, to string) {
	// Parse the CSV file
	transactions, errParse := fp.fileParser.Parse(filePath)
	if errParse != nil {
		log.Fatalf("failed to parse the CSV file: %v", errParse)
	}

	// Generate the summary
	summary := fp.summaryGenerator.Generate(transactions)

	// Send the summary email
	errSend := fp.emailSender.Send(summary.GetStyledEmail(), to)
	if errSend != nil {
		log.Fatalf("failed to send the summary email: %v", errSend)
	}

	fmt.Println("Summary email sent successfully!")

	// Connect to the PostgreSQL database
	db, errOpen := sql.Open("postgres", "host=db port=5432 user=postgres password=postgres dbname=transactions sslmode=disable")
	if errOpen != nil {
		log.Fatalf("failed to connect to the database: %v", errOpen)
	}

	defer db.Close()

	// Save transactions to the database
	errCreate := fp.transacionsRepository.CreateTransactions(transactions)
	if errCreate != nil {
		log.Fatalf("failed to create transactions in the database: %v", errCreate)
	}

	// Get transactions from the database
	dbTransactions, errGet := fp.transacionsRepository.GetAllTransactions()
	if errGet != nil {
		log.Fatalf("failed to get all transactions from the database: %v", errGet)
	}

	// Print the transactions
	printTransactions(dbTransactions)
}

func printTransactions(transactions []domain.Transaction) {
	fmt.Println("Transactions:")
	for _, t := range transactions {
		fmt.Printf("ID: %d\n", t.ID)
		fmt.Printf("Date: %s\n", t.Date.Format("2006-01-02"))
		fmt.Printf("Amount: %.2f\n", t.Amount)
		fmt.Printf("IsCredit: %v\n", t.IsCredit)
		fmt.Printf("IsDebit: %v\n", t.IsDebit)
		fmt.Println()
	}
}
