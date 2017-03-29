package mailgun

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/utils"
	"github.com/rai-project/vipertags"
)

type mailgunConfig struct {
	Provider  string        `json:"provider" config:"email.provider"`
	Source    string        `json:"source" config:"email.source" default:"postmaster@rai-project.com" env:"EMAIL_ADDR"`
	Domain    string        `json:"domain" config:"email.domain" env:"EMAIL_DOMAIN"`
	ApiKey    string        `json:"mailgun_active_api_key" config:"email.mailgun_active_api_key" env:"MAILGUN_API_KEY"`
	PublicKey string        `json:"mailgun_email_validation_key" config:"email.mailgun_email_validation_key" env:"MAILGUN_PUBLIC_API_KEY"`
	done      chan struct{} `json:"-" config:"-"`
}

var (
	Config = &mailgunConfig{
		done: make(chan struct{}),
	}
)

func decrypt(s string) string {
	if utils.IsEncryptedString(s) {
		c, err := utils.DecryptStringBase64(config.App.Secret, s)
		if err == nil {
			return c
		}
	}
	return s
}

func (mailgunConfig) ConfigName() string {
	return "Mailgun"
}

func (a *mailgunConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *mailgunConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	a.ApiKey = decrypt(a.ApiKey)
	a.PublicKey = decrypt(a.PublicKey)
}

func (c mailgunConfig) Wait() {
	<-c.done
}

func (c mailgunConfig) String() string {
	return pp.Sprintln(c)
}

func (c mailgunConfig) Debug() {
	log.Debug("Mailgun Config = ", c)
}

func init() {
	config.Register(Config)
}
