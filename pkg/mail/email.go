package mail

import (
	"fmt"
	"net/smtp"

	"github.com/sirjager/go_emails/cfg"
	"github.com/sirjager/go_emails/pkg/validator"
)

type SMTP struct {
	Address   string
	PlainAuth smtp.Auth
	Email     string
}

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewSMTP(config cfg.Config) (*SMTP, error) {
	err := validator.ValidateEmail(config.SmtpEmail)
	if err != nil {
		return nil, err
	}
	err = validator.ValidatePassword(config.SmtpPass)
	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s:%s", config.SmtpHost, config.SmtpPort)
	plainAuth := smtp.PlainAuth("", config.SmtpEmail, config.SmtpPass, config.SmtpHost)
	return &SMTP{Address: address, PlainAuth: plainAuth, Email: config.SmtpEmail}, nil
}

func (s SMTP) SendMail(mail Mail) (err error) {
	if err = validator.ValidateEmail(mail.To); err != nil {
		return err
	}
	subject := "Subject: " + mail.Subject + " \n"
	message := []byte(subject + mail.Body)
	err = smtp.SendMail(s.Address, s.PlainAuth, mail.From, []string{mail.To}, message)
	return
}
