package adapters

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string, attachFiles []string) error
}
