package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"gin-api-test/config"
	"gin-api-test/db/migrations"
	"gin-api-test/helpers/log"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/fs"
	"path/filepath"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *migrations.User, data *EmailData) {

	from := config.MyEnv.EmailFrom
	smtpHost := config.MyEnv.SmtpHost
	smtpUser := config.MyEnv.SmtpUser
	smtpPass := config.MyEnv.SmtpPassword
	smtpPort := config.MyEnv.SmtpPort

	var body bytes.Buffer
	templates, err := ParseTemplateDir("templates")
	if err != nil {
		log.Log.Error("could not load template", err)
	}
	templates.ExecuteTemplate(&body, "verificationCode.html", &data)
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	fmt.Println(m)
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Log.Error("Could not send email", err)
	}
}
