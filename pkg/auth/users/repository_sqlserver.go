package users

import (
	"database/sql"
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func NewUserSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

func (s *sqlserver) Create(m *User) error  {
	const sqlQueryCreateUser = `insert into auth.user (id, dni, username, name, lastname, phone, email, password, created_at, updated_at) values (:id, :dni, :username, :name, :lastname, :cellphone, :email, :password, :created_at, :updated_at`
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

func (s *sqlserver) Update(m *User) error {
	//TODO implements query updateUser for sqlserver
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

func (s *sqlserver) getByEmail(email string) (*User, error)  {
	const QueryGetUserByEmail = `select id, username, name, lastname, email, password, dni, phone as cellphone, from auth.user where email = @email`
	mdl := User{}
	err := s.DB.Get(&mdl,QueryGetUserByEmail, sql.Named("email",email))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Println(s.TxID, " - couldn't execute QueryGetUserByEmail: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

func (s *sqlserver) getByCellphone(cellphone string) (*User, error)  {
	const QueryGetUserByCellphone = `select id, username, name, lastname, email, password, dni, phone as cellphone, from auth.user where phone = @phone`
	mdl := User{}
	err := s.DB.Get(&mdl,QueryGetUserByCellphone, sql.Named("phone",cellphone))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Println(s.TxID, " - couldn't execute QueryGetUserByCellphone: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}
