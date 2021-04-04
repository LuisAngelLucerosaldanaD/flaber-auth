package users

import (
	"flaber-auth/internal/models"
	"github.com/asaskevich/govalidator"
)

type User models.User

func NewUser(id, name, lastName, email, dni, cellphone string) *User {
	return &User{
		ID: id,
		Name: name,
		LastName: lastName,
		EmailNotifications: email,
		IdentificationNumber: dni,
		Cellphone: cellphone,
	}
}

func (m *User) valid() (bool, error)  {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return false, err
	}
	return result, err
}

