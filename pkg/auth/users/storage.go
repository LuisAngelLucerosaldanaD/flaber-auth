package users

import (
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
)

type ServicesUserRepository interface {
	Create(m *User) error
	Update(m *User) error
	getByEmail(email string) (*User, error)
	getByCellphone(cellphone string) (*User, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUserRepository {
	var s ServicesUserRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return NewUserSqlServerRepository(db, user, txID)
	case Postgresql:
		return NewUserPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
