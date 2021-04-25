package notificacion_email

import (
	"flaber-auth/internal/mail"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	mail *mail.Mail
	TxID string
}

func NewEmailSqlServerRepository(db *sqlx.DB, mail *mail.Mail, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		mail: mail,
		TxID: txID,
	}
}
