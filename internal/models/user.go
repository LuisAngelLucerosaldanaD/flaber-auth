package models

import "time"

type User struct {
	ID                     string     `json:"id" db:"id" valid:"required,uuid"`
	Name                   string     `json:"name,omitempty" db:"name" valid:"required"`
	LastName               string     `json:"lastname,omitempty" db:"lastname" valid:"required"`
	Password               string     `json:"password,omitempty" db:"password" valid:"-"`
	EmailNotifications     string     `json:"email,omitempty" db:"email" valid:"required,email"`
	IdentificationNumber   string     `json:"dni,omitempty" db:"dni" valid:"required"`
	Status                 int        `json:"status,omitempty" db:"status" valid:"-"`
	LastChangePassword     *time.Time `json:"last_change_password,omitempty" db:"last_change_password" valid:"-"`
	BlockDate              *time.Time `json:"block_date,omitempty" db:"block_date" valid:"-"`
	DisabledDate           *time.Time `json:"disabled_date,omitempty" db:"disabled_date" valid:"-"`
	ChangePassword         *bool      `json:"change_password,omitempty" db:"change_password" valid:"-"`
	ChangePasswordDaysLeft *int       `json:"change_password_days_left,omitempty" db:"change_password_days_left" valid:"-"`
	IsBlock                *bool      `json:"is_block,omitempty" db:"is_block" valid:"-"`
	IsDisabled             *bool      `json:"is_disabled,omitempty" db:"is_disabled" valid:"-"`
	UserID 				   string	  `json:"user_id" db:"user_id" valid:"uuid"`
	LastLogin              *time.Time `json:"last_login,omitempty" db:"last_login" valid:"-"`
	TimeOut                int        `json:"time_out,omitempty" valid:"-"`
	ClientID               int        `json:"client_id,omitempty" bson:"client_id"`
	HostName               string     `json:"host_name,omitempty" bson:"host_name"`
	RealIP                 string     `json:"real_ip,omitempty" bson:"real_ip"`
	Token                  string     `json:"token,omitempty" bson:"token"`
	SessionID              string     `json:"session_id" bson:"session_id"`
	Colors                 Color      `json:"colors" bson:"colors"`
	Cellphone              string     `json:"cellphone" bson:"phone" valid:"-"`
	Roles                  []*string  `json:"roles,omitempty" bson:"roles"`
	CreatedAt              time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at,omitempty" db:"updated_at"`
}

type LoggedUsers struct {
	Event     string    `json:"event" bson:"event"`
	HostName  string    `json:"host_name" bson:"host_name"`
	IpRemote  string    `json:"ip_remote" bson:"ip_remote"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type PasswordHistory struct {
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type Color struct {
	Primary   string `json:"primary" bson:"primary"`
	Secondary string `json:"secondary" bson:"secondary"`
	Tertiary  string `json:"tertiary" bson:"tertiary"`
}
