package notificacion_email

import (
	"bytes"
	"flaber-auth/internal/env"
	"flaber-auth/internal/logger"
	"flaber-auth/internal/mail"
	"fmt"
	"github.com/jmoiron/sqlx"
	"html/template"
)

type PortServiceEmail interface {
	SendMail(email string, userCode int64) error
}

type Service struct {
	repository ServicesEmailRepository
	db *sqlx.DB
}

func NewEmailService(repository ServicesEmailRepository, db *sqlx.DB) PortServiceEmail {
	return &Service{ repository: repository ,db: db}
}

func (s *Service) SendMail(email string, userCode int64) error {

	var parameters = make(map[string]string,0)
	e := env.NewConfiguration()
	parameters["TEMPLATE-PATH"] = "notificar_usuario_recuperar_contraseña.gohtml"
	parameters["user_code"] = fmt.Sprintf("%d", userCode)

	body, err := s.generateTemplateMail(parameters)

	if err != nil {
		logger.Error.Printf("couldn't generate body in notification email")
		return err
	}
	var tos []string
	tos = append(tos, email)
	myMail := &mail.Mail{
		From:        e.Smtp.Email,
		To:          tos,
		Subject:     "recovery password",
		Body:        body,
	}

	err = myMail.SendMail()

	if err != nil {
		logger.Error.Printf("couldn't send mail NotificationMail: %V", err)
		return err
	}
	return nil
}

func (s *Service) generateTemplateMail(param map[string]string) (string,error) {
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