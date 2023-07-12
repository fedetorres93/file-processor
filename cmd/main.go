package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/fedtorres/file-processor/internal/adapters"
	"github.com/fedtorres/file-processor/internal/core/services"
	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("please provide the path to the CSV file and the recipient email.")
	}

	filePath := os.Args[1]
	to := os.Args[2]

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	db, errOpen := sql.Open(driverName, fmt.Sprintf(dataSourceName, dbHost, dbPort, dbUser, dbPassword, dbName))
	if errOpen != nil {
		log.Fatalf("failed to connect to the database: %v", errOpen)
	}

	defer db.Close()

	fileParser := adapters.NewFileParser()
	summaryGenerator := adapters.NewSummaryGenerator()
	emailSender := adapters.NewEmailSender(emailFrom, emailPassword)
	transactionsRepository := adapters.NewTransactionsRepository(db)
	fileProcessor := services.NewFileProcessor(fileParser, summaryGenerator, emailSender, transactionsRepository)

	fileProcessor.Process(filePath, to)
}
