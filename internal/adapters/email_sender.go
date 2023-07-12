package adapters

import (
	"fmt"
	"net/smtp"

	"github.com/fedtorres/file-processor/internal/core/domain"
)

const (
	addr = "smtp.gmail.com:587"
)

type emailSender struct {
	from     string
	password string
}

func NewEmailSender(from, password string) *emailSender {
	return &emailSender{
		from:     from,
		password: password,
	}
}

func (es emailSender) Send(msg string, to string) error {
	message := "Subject: Summary of Account Transactions\n"
	message += "Content-Type: text/html; charset=UTF-8\n\n"
	message += msg
	auth := smtp.PlainAuth("", es.from, es.password, "smtp.gmail.com")

	errSend := smtp.SendMail(addr, auth, es.from, []string{to}, []byte(message))
	if errSend != nil {
		fmt.Printf("error while sending email: %s\n", errSend.Error())

		return domain.ErrSendEmail
	}

	return nil
}
