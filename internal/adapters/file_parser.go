package adapters

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fedtorres/file-processor/internal/core/domain"
)

type fileParser struct{}

func NewFileParser() *fileParser {
	return &fileParser{}
}

func (fp fileParser) Parse(filePath string) ([]domain.Transaction, error) {
	file, errOpen := os.Open(filePath)
	if errOpen != nil {
		fmt.Printf("error while opening file: %s\n", errOpen.Error())

		return nil, domain.ErrOpenFile
	}

	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	records, errRead := reader.ReadAll()
	if errRead != nil {
		return nil, domain.ErrReadFile
	}

	var transactions []domain.Transaction

	for i, record := range records {
		if i == 0 {
			continue
		}

		id, errParseID := strconv.Atoi(record[0])
		if errParseID != nil {
			fmt.Printf("failed to parse ID at line %d\n", i+1)

			return nil, domain.ErrParseFile
		}

		date, errParseDate := time.Parse("1/2", record[1])
		if errParseDate != nil {
			fmt.Printf("failed to parse date at line %d\n", i+1)

			return nil, domain.ErrParseFile
		}

		amount, errParseAmount := strconv.ParseFloat(strings.Trim(record[2], " "), 64)
		if errParseAmount != nil {
			fmt.Printf("failed to parse amount at line %d\n", i+1)

			return nil, domain.ErrParseFile
		}

		transaction := domain.Transaction{
			ID:       id,
			Date:     date,
			Amount:   amount,
			IsCredit: strings.HasPrefix(record[2], "+"),
			IsDebit:  strings.HasPrefix(record[2], "-"),
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
