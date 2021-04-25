package notificacion_email

import (
	"flaber-auth/internal/mail"
	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *mail.Mail
	TxID string
}

func NewEmailPsqlRepository(db *sqlx.DB, user *mail.Mail, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}
