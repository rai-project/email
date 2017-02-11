package mailgun

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type mailgunConfig struct {
	Provider     string `json:"provider" config:"email.provider" default:"mailgun"`
	Source       string `json:"source" config:"email.source" default:"postmaster@rai-project.com" env:"EMAIL_ADDR"`
	Domain       string `json:"domain" config:"email.domain" env:"EMAIL_DOMAIN"`
	ApiKey       string `json:"mailgun_api_key" config:"email.mailgun_api_key" env:"MAILGUN_API_KEY"`
	PublicApiKey string `json:"mailgun_public_api_key" config:"email.mailgun_public_api_key" env:"MAILGUN_PUBLIC_API_KEY"`
}

var (
	Config = &mailgunConfig{}
)

func (mailgunConfig) ConfigName() string {
	return "Mailgun"
}

func (mailgunConfig) SetDefaults() {
}

func (a *mailgunConfig) Read() {
	vipertags.Fill(a)
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
