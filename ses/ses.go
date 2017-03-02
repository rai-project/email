package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/rai-project/email"
)

type sesEmail struct {
	*session.Session
}

func New(sess *session.Session) (email.Email, error) {
	if sess != nil {
		return &sesEmail{sess}, nil
	}

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return &sesEmail{sess}, nil
}

func (e *sesEmail) Send(toAddress, subject, body string) error {
	svc := ses.New(e)
	params := &ses.SendEmailInput{
		Destination: &ses.Destination{ // Required
			ToAddresses: []*string{
				aws.String(toAddress), // Required
			},
		},
		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Html: &ses.Content{
					Data:    aws.String(body), // Required
					Charset: aws.String("Charset"),
				},
				Text: &ses.Content{
					Data:    aws.String(body), // Required
					Charset: aws.String("Charset"),
				},
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String(subject), // Required
				Charset: aws.String("Charset"),
			},
		},
		Source: aws.String(Config.Source), // Required
		ReplyToAddresses: []*string{
			aws.String(Config.Source), // Required
		},
		ReturnPath: aws.String(Config.Source),
	}
	_, err := svc.SendEmail(params)
	if err != nil {
		log.WithError(err).
			WithField("to", toAddress).
			Error("Failed to send email message")

		return err
	}

	return nil
}
