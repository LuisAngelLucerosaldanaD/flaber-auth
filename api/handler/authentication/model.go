package authentication

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	NickName  string    `json:"nickname"`
	Email     string    `json:"email"`
	Cellphone string    `json:"cellphone"`
	Password  string    `json:"password"`
	RealIP    string    `json:"real_ip"`
}
