package auth

import (
	"flaber-auth/internal/models"
	"flaber-auth/pkg/auth/login"
	"flaber-auth/pkg/auth/users"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvLogin login.PortServiceLogin
	SrvUsers users.PortServerUser
}

func NewServerAuth(db *sqlx.DB, user *models.User, txID string) *Server  {
	srvLogin := login.NewLoginService(db, txID)

	repoUsers := users.FactoryStorage(db, user, txID)
	srvUsers := users.NewUserService(repoUsers, user, txID)
	return &Server{
		SrvLogin: srvLogin,
		SrvUsers: srvUsers,
	}
}