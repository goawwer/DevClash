package mail

import (
	"os/exec"
	"strings"

	"github.com/goawwer/devclash/pkg/logger"
	"gopkg.in/gomail.v2"
)

type Config struct {
	FromAccount   string `env:"MAIL_FROM_ACCOUNT"`
	AppPassword   string `env:"MAIL_APP_PASSWORD"`
	TemplatesPath string `env:"MAIL_TEMPLATES_PATH"`
}

var d *gomail.Dialer
var msg *gomail.Message
var templatesPath string

func Init(cfg *Config) {
	dealer := gomail.NewDialer("smtp.gmail.com", 587, cfg.FromAccount, cfg.AppPassword)
	message := gomail.NewMessage()

	output, err := exec.Command("go", "list", "-f", "{{.Module.Dir}}").Output()
	if err != nil {
		logger.Error("failed to take output from exec.Command(pwd)")
	}

	templatesPath = strings.TrimSpace(string(output)) + cfg.TemplatesPath
	d = dealer
	msg = message
}
