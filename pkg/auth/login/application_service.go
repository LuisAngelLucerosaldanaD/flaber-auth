package login

import (
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	password2 "flaber-auth/internal/password"
	"flaber-auth/pkg/auth/users"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PortServiceLogin interface {
	Login(email, cellphone, password, ralIp string) (string, int, error)
}

type Service struct {
	DB   *sqlx.DB
	TxID string
}

func NewLoginService(db *sqlx.DB, TxID string) PortServiceLogin {
	return &Service{DB: db, TxID: TxID}
}

func (s *Service) Login(email, cellphone, password, realIP string) (string, int, error) {
	var token string
	m := NewLogin(password, cellphone, email, realIP)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.TxID, " - don't meet validations: ", err)
		return "", 15, err
	}
	usr, code, err := s.getUser(email, cellphone)
	if err != nil {
		logger.Error.Println(s.TxID, " - couldn't get user by id: ", err)
		return token, code, err
	}
	if !password2.Compare(usr.ID, usr.Password, m.Password) {
		return token, 10, fmt.Errorf("usuario o contraseña incorrectos")
	}
	usr.RealIP = m.RealIP
	usr.Password = ""
	token, code, err = GenerateJWT(usr)
	if err != nil {
		logger.Error.Println(s.TxID, " - couldn't generate token:%s - %v", s.TxID, err)
		return "", code, err
	}
	return token, 29, nil
}

func (s *Service) getUser(email, cellphone string) (*models.User, int, error) {
	repoLogin := users.FactoryStorage(s.DB, nil, s.TxID)
	srvLogin := users.NewUserService(repoLogin, nil, s.TxID)
	if email != "" {
		user, _, err := srvLogin.GetUserByEmail(email)
		if err != nil {
			logger.Error.Println("Couldn't get user by email", err)
			return nil, 10, err
		}
		if user == nil {
			logger.Error.Println("Couldn't get user by nickName")
			return nil, 10, fmt.Errorf("couldn't get user by email")
		}
		usr := models.User(*user)
		return &usr, 29, nil
	}
	if cellphone != "" {
		user, _, err := srvLogin.GetUserByCellphone(cellphone)
		if err != nil {
			logger.Error.Println("couldn't get user by cellphone")
			return nil, 10, err
		}
		if user == nil{
			logger.Error.Println("couldn't get user by cellphone")
			return nil, 10, fmt.Errorf("couldn't get user by email")
		}
		usr := models.User(*user)
		return &usr, 29, nil
	}
	return nil, 1, fmt.Errorf("user don't exists")
}