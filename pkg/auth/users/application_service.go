package users

import (
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"flaber-auth/internal/password"
	"fmt"
)

type PortServerUser interface {
	CreateUser(id, name, lastName, Password, email, identificationNumber, cellphone string) (*User, int, error)
	GetUserByEmail(email string) (*User, int, error)
	GetUserByCellphone(cellphone string) (*User, int, error)
}

type Service struct {
	repository ServicesUserRepository
	user       *models.User
	txID       string
}

func NewUserService(repository ServicesUserRepository, user *models.User, TxID string) PortServerUser {
	return &Service{repository: repository, user: user, txID: TxID}
}

func (s *Service) CreateUser(id, name, lastName, Password, email, identificationNumber, cellphone string) (*User, int, error) {
	var changePass, isBlock, isDisabled bool
	m := NewUser(id, name, lastName, email, identificationNumber, cellphone)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't create user: ", err)
		return nil, 15, err
	}
	user, _, err := s.GetUserByCellphone(m.Cellphone)
	if err != nil{
		logger.Error.Println(s.txID, " - don't create user: ", err)
		return nil, 15, err
	}
	if user != nil {
		logger.Error.Println(s.txID, " - the user could not be created, this phone is already being used by another user.")
		return nil, 23, fmt.Errorf("the user could not be created, this phone is already being used by another user")
	}
	user, _, err = s.GetUserByEmail(m.EmailNotifications)
	if err != nil{
		logger.Error.Println(s.txID, " - don't create user: ", err)
		return nil, 15, err
	}
	if user != nil {
		logger.Error.Println(s.txID, " - the user could not be created, this email is already being used by another user.")
		return nil, 22, fmt.Errorf("the user could not be created, this email is already being used by another user")
	}
	m.Password = password.Encrypt(Password)
	if m.ChangePassword == nil {
		m.ChangePassword = &changePass
	}
	m.IsBlock = &isBlock
	m.IsDisabled = &isDisabled
	if err := s.repository.Create(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't create User: ", err)
		if err.Error() == "Fluber: 108" {
			return nil, 108, err
		}
		return nil, 3, err
	}
	return m, 29, nil
}

func (s *Service) GetUserByEmail(email string) (*User, int, error) {
	m, err := s.repository.getByEmail(email)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getUserByEmail row: ", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *Service) GetUserByCellphone(cellphone string) (*User, int, error) {
	m, err := s.repository.getByCellphone(cellphone)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByCellphone row: ", err)
		return nil, 22,err
	}
	return m, 29, nil
}