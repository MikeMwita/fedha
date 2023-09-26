package mailing

import (
	"github.com/MikeMwita/fedha.git/services/app-notification/config"
	"github.com/stretchr/testify/require"
	"testing"
)

//
//func TestNewSendMail(t *testing.T) {
//	conf, err := config.LoadConfig()
//	require.NoError(t, err)
//
//	s := NewSendMail(conf.Email.Name, conf.Email.FromEmailAddr, conf.Email.FromEmailPassword)
//
//	require.NoError(t, err)
//
//	subject := "A test email"
//	content := `
//	<h1>Hello world</h1>
//	<p>This is a test message from <a href="http://fedha.co.ke">Fedha</a></p>
//	`
//	to := []string{"fakemike285@gmail.com"}
//	attachFiles := []string{"./Readme.md"}
//
//	err = s.SendEmail(subject, content, to, nil, attachFiles)
//	require.NoError(t, err)
//}

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
    <p>This is a test message from <a href="http://techschool.guru">Tech School</a></p>
    `
	to := []string{"fakemike285@gmail.com"}
	attachFiles := []string{"./README.md"}

	err = sender.SendEmail(subject, content, to, nil, attachFiles) // Remove the extra 'nil' argument
	require.NoError(t, err)
}

err = sender.SendEmail(subject, content, to, nil, attachFiles)
if err != nil {
fmt.Printf("Error sending email: %v\n", err)
fmt.Printf("Subject: %s\n", subject)
fmt.Printf("Content:\n%s\n", content)
fmt.Printf("To: %v\n", to)
fmt.Printf("AttachFiles: %v\n", attachFiles)
}
