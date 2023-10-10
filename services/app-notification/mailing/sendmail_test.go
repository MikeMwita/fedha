package mailing

import (
	"github.com/MikeMwita/fedha.git/services/app-notification/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	conf, err := config.LoadConfig("..")
	require.NoError(t, err)

	sender := NewSendMail(conf.EmailSenderName, conf.EmailSenderAddress, conf.EmailSenderPassword)
	subject := "A test email"
	content := `
    <h1>Hello world</h1>
    <p>This is a test message from <a href="http://fedha.guru">Fedha</a></p>
    `
	to := []string{"fakemike285@gmail.com"}
	attachFiles := []string{"./README.md"}

	err = sender.SendEmail(subject, content, to, nil, attachFiles) // Remove the extra 'nil' argument
	require.NoError(t, err)
}
