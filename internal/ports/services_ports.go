package ports

import "github.com/fedtorres/file-processor/internal/core/domain"

type (
	FileParserPort interface {
		Parse(filePath string) ([]domain.Transaction, error)
	}

	SummaryGeneratorPort interface {
		Generate(transactions []domain.Transaction) domain.Summary
	}

	EmailSenderPort interface {
		Send(msg, to string) error
	}
)
