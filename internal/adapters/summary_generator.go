package adapters

import (
	"github.com/fedtorres/file-processor/internal/core/domain"
)

type summaryGenerator struct{}

func NewSummaryGenerator() *summaryGenerator {
	return &summaryGenerator{}
}

func (sg summaryGenerator) Generate(transactions []domain.Transaction) domain.Summary {
	summary := domain.Summary{
		TotalTransactionsCounts:  map[string]int{},
		DebitTransactionsCounts:  map[string]int{},
		CreditTransactionsCounts: map[string]int{},
		AverageDebitAmounts:      map[string]float64{},
		AverageCreditAmounts:     map[string]float64{},
	}

	for _, transaction := range transactions {
		summary.TotalBalance += transaction.Amount

		month := transaction.Date.Month().String()
		summary.TotalTransactionsCounts[month]++

		if transaction.IsDebit {
			summary.DebitTransactionsCounts[month]++
			summary.AverageDebitAmounts[month] += transaction.Amount
		}

		if transaction.IsCredit {
			summary.CreditTransactionsCounts[month]++
			summary.AverageCreditAmounts[month] += transaction.Amount
		}
	}

	for month, count := range summary.DebitTransactionsCounts {
		summary.AverageDebitAmounts[month] = summary.AverageDebitAmounts[month] / float64(count)
	}

	for month, count := range summary.CreditTransactionsCounts {
		summary.AverageCreditAmounts[month] = summary.AverageCreditAmounts[month] / float64(count)
	}

	return summary
}
