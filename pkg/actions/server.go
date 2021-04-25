package actions

import (
	"flaber-auth/internal/mail"
	"flaber-auth/pkg/actions/notificacion_email"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvMail notificacion_email.PortServiceEmail
}

func NewServerMail(db *sqlx.DB, mail *mail.Mail, txID string) *Server  {

	repoUsers := notificacion_email.FactoryStorage(db, mail, txID)
	srvMails := notificacion_email.NewEmailService(repoUsers, db)

	return &Server{
		SrvMail: srvMails,
	}
}
