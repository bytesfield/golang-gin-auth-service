package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunSend struct {
	sender       string
	recipient    string
	mailSubject  string
	mailTemplate string
}

func (ms *MailgunSend) To(email string) *MailgunSend {
	ms.sender = email
	return ms
}

func (ms *MailgunSend) From(email string) *MailgunSend {
	ms.recipient = email
	return ms
}

func (ms *MailgunSend) Subject(subject string) *MailgunSend {
	ms.mailSubject = subject
	return ms
}

func (ms *MailgunSend) Template(template string) *MailgunSend {
	ms.mailTemplate = template
	return ms
}

func (ms *MailgunSend) Send() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(ms.sender, ms.mailSubject, ms.mailTemplate, ms.recipient)

	// message.SetTemplate(ms.mailTemplate)
	// err := message.AddTemplateVariable("passwordResetLink", "some link to your site unique to your user")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
