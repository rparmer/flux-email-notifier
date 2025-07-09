package email

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v5"
	"github.com/rparmer/flux-email-notifier/config"
)

type Email struct {
	Message string
	Subject string
	From    Contact
	To      Contact
}

type Contact struct {
	Name    string
	Address string
}

func New() *Email {
	var e Email
	return &e
}

func (e *Email) Send() error {
	cfg := config.GetConfig()
	to := e.To.Address
	from := e.From.Address

	if e.From.Name != "" {
		from = fmt.Sprintf("%s <%s>", e.From.Name, e.From.Address)
	}

	message := mailgun.NewMessage(cfg.Mailgun.Domain, from, e.Subject, e.Message, to)
	mg := mailgun.NewMailgun(cfg.Mailgun.Key)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := mg.Send(ctx, message)
	if err != nil {
		return err
	}
	if response.Message != "" {
		fmt.Println(response.Message)
	}
	return nil
}
