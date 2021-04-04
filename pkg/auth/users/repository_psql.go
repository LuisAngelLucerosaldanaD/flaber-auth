package users

import (
	"database/sql"
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func NewUserPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

func (s *psql) Create(m *User) error  {
	date := time.Now()
	m.CreatedAt = date
	m.UpdatedAt = date
	const sqlQueryCreateUser = `insert into auth.user (id, dni, name, lastname, phone, email, password, created_at, updated_at) values (:id, :dni, :name, :lastname, :cellphone, :email, :password, :created_at, :updated_at)`
	rs, err := s.DB.NamedExec(sqlQueryCreateUser, &m)
	if err != nil {
		logger.Error.Println(s.TxID, " - couldn't insert User: %V", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Fluber: 108")
	}
	return nil
}

func (s *psql) Update(m *User) error {
	//TODO implements query updateUser for psql
	const sqlQueryUpdateUser = ``
	rs, err := s.DB.NamedExec(sqlQueryUpdateUser, &m)
	if err != nil {
		logger.Error.Println(s.TxID, " - couldn't update User: %V", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Fluber: 108")
	}
	return nil
}

func (s *psql) getByEmail(email string) (*User, error)  {
	const QueryGetUserByEmail = `select id, username, name, lastname, email, password, dni, phone as cellphone, from auth.user where email = $1`
	mdl := User{}
	err := s.DB.Get(&mdl,QueryGetUserByEmail, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Println(s.TxID, " - couldn't execute QueryGetUserByEmail: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

func (s *psql) getByCellphone(cellphone string) (*User, error)  {
	const QueryGetUserByCellphone = `select id, username, name, lastname, email, password, dni, phone as cellphone from auth.user where phone = $1`
	mdl := User{}
	err := s.DB.Get(&mdl,QueryGetUserByCellphone, cellphone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Println(s.TxID, " - couldn't execute QueryGetUserByCellphone: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}