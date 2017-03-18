package ses

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type sesConfig struct {
	Provider string        `json:"provider" config:"email.provider"`
	Source   string        `json:"source" config:"email.source" default:"postmaster@rai-project.com" env:"EMAIL_ADDR"`
	Domain   string        `json:"domain" config:"email.domain" env:"EMAIL_DOMAIN"`
	done     chan struct{} `json:"-" config:"-"`
}

var (
	Config = &sesConfig{
		done: make(chan struct{}),
	}
)

func (*sesConfig) ConfigName() string {
	return "SES"
}

func (a *sesConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *sesConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
}

func (c sesConfig) Wait() {
	<-c.done
}

func (c *sesConfig) String() string {
	return pp.Sprintln(c)
}

func (c *sesConfig) Debug() {
	log.Debug("SES Config = ", c)
}

func init() {
	config.Register(Config)
}
