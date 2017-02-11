package email

type Email interface {
	Send(toAddress, subject, body string) error
}
