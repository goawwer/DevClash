package mail

import (
	"bytes"
	"html/template"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	organizermodel "github.com/goawwer/devclash/internal/domain/organizer_model"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/sirupsen/logrus"
)

func sendMessage(toEmail string, subject string, body *bytes.Buffer) error {
	msg.SetHeader("From", d.Username)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body.String())

	return d.DialAndSend(msg)
}

func OrganizerSignupInfo(a *accountmodel.Account, org *organizermodel.OrganizerAccount) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatesPath + "/organizer_signup.html")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"component": "mail service",
			"info":      "failed to parse sign up organizer html document",
		})
		return err
	}
	t.Execute(&body, struct {
		Name string
	}{
		Name: org.Name,
	})

	return sendMessage(a.Email, "Регистрация", &body)
}
