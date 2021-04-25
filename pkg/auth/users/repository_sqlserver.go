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
	const sqlQueryCreateUser = `insert into auth.user (id, dni name, lastname, phone, email, password, created_at, updated_at) values (:id, :dni, :name, :lastname, :cellphone, :email, :password, :created_at, :updated_at`
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

func (s *sqlserver) UpdatePasswordByUserId(password, userId string) error {
	m := User{Password: password, ID: userId}
	const sqlQueryUpdatePasswordByUserId = `UPDATE auth."user" SET password = :password where id = :id;`
	rs, err := s.DB.NamedExec(sqlQueryUpdatePasswordByUserId, &m)
	if err != nil {
		logger.Error.Println(s.TxID, " - couldn't update password by user id: %V", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Fluber: 108")
	}
	return nil
}

func (s *sqlserver) UpdateCodeByEmailAndUserID(code int64, email, userID string) error {
	m := User{UserCode: code,EmailNotifications: email, ID: userID}
	const SqlUpdateCodeByEmailAndUserID = `UPDATE auth."user" SET  user_code= :user_code WHERE (email = :email) and (id = :id);`
	rs, err := s.DB.NamedExec(SqlUpdateCodeByEmailAndUserID, &m)
	if err != nil {
		logger.Error.Println(s.TxID, " - couldn't update User code by email and userId: %V", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("Fluber: 108")
	}
	return nil
}

func (s *sqlserver) GetUserById(userId string) (*User, error) {
	const SqlQueryGetUserById = `SELECT id, name, lastname, phone, email, password, created_at, updated_at, user_code, is_deleted
	FROM auth."user" where id = @userID;`
	mdl := User{}
	err := s.DB.Get(&mdl, SqlQueryGetUserById, sql.Named("userID", userId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Println(s.TxID, " - couldn't execute SqlQueryGetUserById: %v", err)
		return &mdl, err
	}
	return &mdl, err
}

func (s *sqlserver) getByEmail(email string) (*User, error)  {
	const QueryGetUserByEmail = `select id, name, lastname, email, password, dni, phone from auth.user where email = @email`
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
	const QueryGetUserByCellphone = `select id, name, lastname, email, password, dni, phone from auth.user where phone = @phone`
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
