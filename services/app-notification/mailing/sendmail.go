package mailing

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	smtpAuthAddr   = "smtp.gmail.com"
	smtpServerAddr = "smtp.gmail.com:587"
)

type SendMail struct {
	name              string
	fromEmailAddr     string
	fromEmailPassword string
}

func (s *SendMail) SendEmail(subject string, content string, to []string, cc []string, attachFiles []string) error {

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.name, s.fromEmailAddr)

	//e.From = s.name + " <" + s.fromEmailAddr + ">"
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc

	attachFiles = []string{"./Readme.md"}

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file: %w", err)
		}
	}
	return e.Send(smtpServerAddr, smtp.PlainAuth("", s.fromEmailAddr, s.fromEmailPassword, smtpAuthAddr))
}

func NewSendMail(name string, fromEmailAddr string, fromEmailPassword string) *SendMail {
	return &SendMail{name: name, fromEmailAddr: fromEmailAddr, fromEmailPassword: fromEmailPassword}
}
