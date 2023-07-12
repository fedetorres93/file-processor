package domain

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID       int
	Date     time.Time
	Amount   float64
	IsDebit  bool
	IsCredit bool
}

type Summary struct {
	TotalBalance             float64
	TotalTransactionsCounts  map[string]int
	DebitTransactionsCounts  map[string]int
	CreditTransactionsCounts map[string]int
	AverageDebitAmounts      map[string]float64
	AverageCreditAmounts     map[string]float64
}

func (s Summary) GetStyledEmail() string {
	summary := "<html><head><style>"
	summary += "body { font-family: Arial, sans-serif; }"
	summary += "h2 { color: #333333; }"
	summary += "table { border-collapse: collapse; width: 100%; margin-bottom: 20px; }"
	summary += "th, td { border: 1px solid #333333; padding: 10px; text-align: left; }"
	summary += "th { background-color: #f2f2f2; }"
	summary += "tr:nth-child(even) { background-color: #f9f9f9; }"
	summary += "tr:hover { background-color: #f5f5f5; }"
	summary += "img { display: block; margin: 20px auto; max-width: 100%; height: auto; }"
	summary += "</style></head><body>"
	summary += "<h2>Summary of Account Transactions</h2>"
	summary += fmt.Sprintf("<p>Total balance: $%.2f</p>", s.TotalBalance)
	summary += "<table>"
	summary += "<tr><th>Month</th><th>Transaction Count</th><th>Average Debit</th><th>Average Credit</th></tr>"

	for month, count := range s.TotalTransactionsCounts {
		summary += "<tr>"
		summary += fmt.Sprintf("<td>%s</td>", month)
		summary += fmt.Sprintf("<td>%d</td>", count)
		summary += fmt.Sprintf("<td>$%.2f</td>", s.AverageDebitAmounts[month])
		summary += fmt.Sprintf("<td>$%.2f</td>", s.AverageCreditAmounts[month])
		summary += "</tr>"
	}

	summary += "</table>"
	summary += "<img src='https://pbs.twimg.com/profile_images/1547301538137341957/LXlcLhCT_400x400.jpg' alt='Transaction Summary'>"
	summary += "</body></html>"

	return summary
}
