package users

import (
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"flaber-auth/internal/password"
	"fmt"
)

type PortServerUser interface {
	CreateUser(id, name, lastName, Password, email, cellphone string, userCode int64) (*User, int, error)
	GetUserByEmail(email string) (*User, int, error)
	GetUserByCellphone(cellphone string) (*User, int, error)
	UpdateCodeByEmailAndUserID(code int64, email, userId string) error
	UpdatePasswordByUserId(pass, userId string) error
	GetUserById(userId string) (*User, error)
}

type Service struct {
	repository ServicesUserRepository
	user       *models.User
	txID       string
}

func NewUserService(repository ServicesUserRepository, user *models.User, TxID string) PortServerUser {
	return &Service{repository: repository, user: user, txID: TxID}
}

func (s *Service) CreateUser(id, name, lastName, Password, email, cellphone string, userCode int64) (*User, int, error) {
	var changePass bool
	m := NewUser(id, name, lastName, email, cellphone, userCode)
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
	if err := s.repository.Create(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't create User: ", err)
		if err.Error() == "Fluber: 108" {
			return nil, 108, err
		}
		return nil, 3, err
	}
	return m, 29, nil
}

func (s *Service) UpdateCodeByEmailAndUserID(code int64, email, userId string) error {
	return s.repository.UpdateCodeByEmailAndUserID(code, email, userId)
}

func (s *Service) UpdatePasswordByUserId(pass, userId string) error {
	user, err := s.GetUserById(userId)
	if err != nil {
		logger.Error.Println("Couldn't get user by id: %V", err)
		return err
	}
	if password.Compare(userId, user.Password, pass) {
		logger.Error.Println("Error, la contraseña nueva no puede ser igual a la contraseña antigua.")
		return fmt.Errorf("Error, la contraseña nueva no puede ser igual a la contraseña antigua. ")
	}

	pass = password.Encrypt(pass)

	return s.repository.UpdatePasswordByUserId(pass, userId)
}

func (s *Service) GetUserById(userId string) (*User, error) {
	return s.repository.GetUserById(userId)
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