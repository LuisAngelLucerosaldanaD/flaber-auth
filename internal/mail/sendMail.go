package mail

import (
	"bytes"
	"crypto/tls"
	"flaber-auth/internal/env"
	"flaber-auth/internal/logger"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"time"
)

func (m *Mail) SendMail() error {
	c := env.NewConfiguration()
	e := gomail.NewMessage()
	e.SetHeader("From", m.From)
	e.SetHeader("To", m.To...)
	e.SetHeader("Cc", m.CC...)
	e.SetHeader("Subject", m.Subject)
	e.SetBody("text/html", m.Body)
	if len(m.Attach) > 0 {
		//m.Attach(e.Attach)
	}

	for _, v := range m.Attachments {
		e.Attach(v)
	}

	mp := c.Smtp.Port
	d := gomail.NewDialer(c.Smtp.Host, mp, c.Smtp.Email, c.Smtp.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(e)
	if err != nil {
		logger.Error.Printf("couldn't emil to: %s, subject: %s, %v", m.To, m.Subject, err)
		return err
	}

	return nil
}

func (m *Mail) AddAttach(fn string) {
	if len(m.Attachments) == 0 {
		m.Attachments = make([]string, 0)
	}
	m.Attachments = append(m.Attachments, fn)
}

func (m *Mail) SendMailNotification(template string, userID string, toMail string, subject string) {
	myMail := Mail{}
	param := make(map[string]string)

	param["TEMPLATE-PATH"] = template
	param["user"] = userID
	param["FECHA-EXECUTE"] = time.Now().String()
	param["TO-MAIL"] = toMail
	param["FROM-EMAIL"] = "luisangellucerosaldana20@gmail.com"
	param["SUBJECT-EMAIL"] = subject

	body, err := m.GenerateTemplateMail(param)
	if err != nil {
		logger.Error.Printf("couldn't generate body in NotificationEmail: %v", err)
		return
	}

	email := param["TO-MAIL"]
	var tos = []string{email}

	myMail.From = param["FROM-EMAIL"]
	myMail.To = tos
	myMail.Subject = fmt.Sprintf(`%s`, param["SUBJECT-EMAIL"])
	myMail.Body = body

	err = myMail.SendMail()
	if err != nil {
		logger.Error.Printf("couldn't sendMail NotificationEmail: %v", err)
		return
	}

	return
}


func (m *Mail) GenerateTemplateMail(param map[string]string) (string, error) {
	bf := &bytes.Buffer{}
	tpl := &template.Template{}

	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
	err := tpl.ExecuteTemplate(bf, param["TEMPLATE-PATH"], &param)
	if err != nil {
		logger.Error.Printf("couldn't generate template body mail: %v", err)
		return "", err
	}
	return bf.String(), err
}
