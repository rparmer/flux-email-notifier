package email

import (
	"fmt"

	"github.com/rparmer/flux-email-notifier/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

func (e *Email) Send() (int, error) {
	cfg := config.GetConfig()
	from := mail.NewEmail(e.From.Name, e.From.Address)
	to := mail.NewEmail(e.To.Name, e.To.Address)
	message := mail.NewSingleEmail(from, e.Subject, to, e.Message, "")
	client := sendgrid.NewSendClient(cfg.Sendgrid.Key)
	response, err := client.Send(message)
	if err != nil {
		return 0, err
	}
	if response.Body != "" {
		fmt.Println(response.Body)
	}
	return response.StatusCode, nil
}
