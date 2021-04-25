package notificacion_email

import (
	"flaber-auth/internal/logger"
	"flaber-auth/internal/mail"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
)

type ServicesEmailRepository interface {

}

func FactoryStorage(db *sqlx.DB, mail *mail.Mail, txID string) ServicesEmailRepository {
	var s ServicesEmailRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return NewEmailSqlServerRepository(db, mail, txID)
	case Postgresql:
		return NewEmailPsqlRepository(db, mail, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}

