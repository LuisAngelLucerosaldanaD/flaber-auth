package login

import (
	"github.com/asaskevich/govalidator"
)

// Model estructura de Module
type Login struct {
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email" valid:"-"`
	Cellphone string `json:"cellphone" db:"phone"`
	RealIP    string `json:"real_ip" db:"real_ip"`
}

func NewLogin(password, cellphone, email, RealIP string) *Login {
	return &Login{
		Password: password,
		Cellphone: cellphone,
		Email:    email,
		RealIP:   RealIP,
	}
}

func (m *Login) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
