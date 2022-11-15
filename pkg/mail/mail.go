package mail

import (
	"sync"

	"github.com/shanedoc/gohub/pkg/config"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // Plaintext message (optional)
	HTML    []byte // Html message (optional)
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

//单例获取邮件对象
func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})
	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, config.GetStringMapString("mail.smtp"))
}
