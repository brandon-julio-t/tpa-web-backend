package facades

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
	"os"
)

type mail struct {
	Client *mailjet.Client
}

var mailSingleton *mail

func getMail() *mail {
	if mailSingleton == nil {
		mailSingleton = &mail{
			Client: mailjet.NewMailjetClient(os.Getenv("MAILJET_PUBLIC_KEY"), os.Getenv("MAILJET_PRIVATE_KEY")),
		}
	}
	return mailSingleton
}

func UseMail() *mail {
	return getMail()
}

func (m mail) SendText(text, subject, id, recipientEmail string) error {
	result, err := m.Client.SendMailV31(
		&mailjet.MessagesV31{
			Info: []mailjet.InfoMessagesV31{
				{
					From: &mailjet.RecipientV31{
						Email: "brandon.julio.t@icloud.com",
						Name:  "STAEM",
					},
					To: &mailjet.RecipientsV31{
						mailjet.RecipientV31{
							Email: recipientEmail,
						},
					},
					Subject:  subject,
					TextPart: text,
					CustomID: id,
				},
			},
		},
	)

	log.Print(result)
	return err
}

func (m mail) SendHTML(html, subject, id, recipientEmail string) error {
	result, err := m.Client.SendMailV31(
		&mailjet.MessagesV31{
			Info: []mailjet.InfoMessagesV31{
				{
					From: &mailjet.RecipientV31{
						Email: "brandon.julio.t@icloud.com",
						Name:  "STAEM",
					},
					To: &mailjet.RecipientsV31{
						mailjet.RecipientV31{
							Email: recipientEmail,
						},
					},
					Subject:  subject,
					HTMLPart: html,
					CustomID: id,
				},
			},
		},
	)

	log.Print(result)
	return err
}
