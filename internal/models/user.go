package models

import "time"

type User struct {
	ID                 string    `json:"id" db:"id" valid:"required,uuid"`
	Name               string    `json:"name,omitempty" db:"name" valid:"required"`
	LastName           string    `json:"lastname,omitempty" db:"lastname" valid:"required"`
	Password           string    `json:"password,omitempty" db:"password" valid:"-"`
	EmailNotifications string    `json:"email,omitempty" db:"email" valid:"required,email"`
	ChangePassword     *bool     `json:"change_password,omitempty" db:"change_password" valid:"-"`
	UserID             string    `json:"user_id" db:"user_id" valid:"uuid"`
	UserCode           int64     `json:"user_code" db:"user_code"`
	HostName           string    `json:"host_name,omitempty" db:"host_name"`
	RealIP             string    `json:"real_ip,omitempty" db:"real_ip"`
	Cellphone          string    `json:"cellphone" db:"phone" valid:"-"`
	CreatedAt          time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt          time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	IsDeleted          bool      `json:"is_deleted,omitempty" db:"is_deleted"`
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
