package mailgun

import (
	mailgun "gopkg.in/mailgun/mailgun-go.v1"

	"github.com/rai-project/email"
)

type mailgunEmail struct {
	mailgun.Mailgun
}

func New() (email.Email, error) {
	mg := mailgun.NewMailgun(Config.Domain, Config.ApiKey, Config.PublicKey)
	return &mailgunEmail{mg}, nil
}

func (e *mailgunEmail) Send(toAddress, subject, body string) error {
	m := e.Mailgun.NewMessage(Config.Source, subject, body, toAddress)
	msg, id, err := e.Mailgun.Send(m)
	if err != nil {
		log.WithError(err).
			WithField("message", msg).
			WithField("id", id).
			WithField("to", toAddress).
			Error("Failed to send email message")
		return err
	}
	return nil
}
