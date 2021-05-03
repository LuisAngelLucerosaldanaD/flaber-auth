package user

import (
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"flaber-auth/pkg/actions"
	"flaber-auth/pkg/auth"
	"flaber-auth/pkg/auth/login"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"net/http"
	"time"
)

type Handler struct {
	DB *sqlx.DB
	TxID string
}
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	m := models.User{}
	err := c.BodyParser(&m)
	if err != nil {
		res.Code = 1
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	srvNxm := auth.NewServerAuth(h.DB, nil, h.TxID)
	usr, cod, err := srvNxm.SrvUsers.CreateUser(m.ID, m.Name, m.LastName, m.Password, m.EmailNotifications, m.Cellphone, 0)
	if err != nil {
		res.Code = cod
		res.Type = "Error"
		res.Msg = "don't create user"
		logger.Error.Println("No se pudo crear el usuario: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	user := models.User{
		ID:                 usr.ID,
		Name:               usr.Name,
		LastName:           usr.LastName,
		EmailNotifications: usr.EmailNotifications,
		UserCode:           usr.UserCode,
		HostName:           c.Hostname(),
		RealIP:             c.IP(),
		Cellphone:          usr.Cellphone,
	}
	token, code, err := login.GenerateJWT(&user)
	if err != nil {
		res.Code = code
		res.Type = "Error"
		res.Msg = "don't create user"
		logger.Error.Println("No se pudo crear el token: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = token
	res.Code = 29
	res.Type = "success"
	res.Msg = "El usuario ha sido creado exitosamente"
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) ChangePassword(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	m := requestChangePassword{}
	err := c.BodyParser(&m)
	if err != nil {
		res.Code = 15
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if m.Password != m.ConfirmPassword {
		res.Code = 2
		res.Type = "Error"
		res.Msg = "Las constraseñas no coinciden, validar por favor."
		logger.Error.Println("Las constraseñas no coinciden, validar por favor.")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvNxm := auth.NewServerAuth(h.DB, nil, h.TxID)
	err = srvNxm.SrvUsers.UpdatePasswordByUserId(m.Password, m.UserID)
	if err != nil {
		res.Code = 3
		res.Type = "Error"
		res.Msg = err.Error()
		logger.Error.Println("No se actualizar la contraseña: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Code = 29
	res.Type = "success"
	res.Msg = "La contraseña se ha cambiado exitosamente"
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) RecoveryPassword(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	m := requestRecoveryPassword{}
	err := c.BodyParser(&m)
	if err != nil {
		res.Code = 15
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvNxm := auth.NewServerAuth(h.DB, nil, h.TxID)
	usr, cod, err := srvNxm.SrvUsers.GetUserByEmail(m.Email)
	if err != nil {
		res.Code = cod
		res.Type = "Error"
		res.Msg = "Error a la hora de obtener el usuario por el email"
		logger.Error.Println("Error a la hora de obtener el usuario por el email : %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if usr == nil {
		res.Code = cod
		res.Type = "Error"
		res.Msg = "El email introducido no corresponde a ninguna cuenta registrada"
		logger.Error.Println("El email introducido no corresponde a ninguna cuenta registrada : %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if usr.UserCode != m.UserCode {
		res.Code = 5
		res.Type = "Error"
		res.Msg = "El codigo de seguridad ingresado no es el correcto"
		logger.Error.Println("El codigo de seguridad ingresado no es el correcto.")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Code = 29
	res.Type = "success"
	res.Msg = "El codigo de seguridad ingresado es correcto"
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) ValidEmailAndSaveCode(c *fiber.Ctx) error {
	min := 1000
	max := 9999
	res := models.Response{Error: true}
	m := requestValidAndChangeCode{}
	var userCode int64
	err := c.BodyParser(&m)
	if err != nil {
		res.Code = 15
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvNxm := auth.NewServerAuth(h.DB, nil, h.TxID)
	usr, cod, err := srvNxm.SrvUsers.GetUserByEmail(m.Email)
	if err != nil {
		res.Code = cod
		res.Type = "Error"
		res.Msg = "Error a la hora de obtener el usuario por el email"
		logger.Error.Println("Error a la hora de obtener el usuario por el email : %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if usr == nil {
		res.Code = cod
		res.Type = "Error"
		res.Msg = "El email introducido no corresponde a ninguna cuenta registrada"
		logger.Error.Println("El email introducido no corresponde a ninguna cuenta registrada : %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	rand.Seed(time.Now().UnixNano())
	userCode = int64(rand.Intn(max - min + 1) + min)

	if userCode == usr.UserCode {
		userCode = int64(rand.Intn(max - min + 1) + min)
	}

	err = srvNxm.SrvUsers.UpdateCodeByEmailAndUserID(userCode, usr.EmailNotifications, usr.ID)
	if err != nil {
		res.Code = 2
		res.Type = "Error"
		res.Msg = "No se pudo actualizar el codigo de recuperarción del usuario"
		logger.Error.Println("No se pudo actualizar el codigo de recuperarción del usuario : %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}


	srvFlb := actions.NewServerMail(h.DB, nil, h.TxID)

	err = srvFlb.SrvMail.SendMail(m.Email, userCode)

	if err != nil {
		res.Code = 13
		res.Type = "Error"
		res.Msg = "error when execute send mail"
		logger.Error.Println("error when execute send mail: %V", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	
	res.Code = 29
	res.Data = usr
	res.Type = "success"
	res.Msg = "Se ha enviado el codigo de recuperación a el email correctament."
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) ExistEmail(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	m := RequestExistEmail{}
	err := c.BodyParser(&m)
	if err != nil {
		res.Code = 1
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvNxm := auth.NewServerAuth(h.DB, nil, h.TxID)
	usr, _, err := srvNxm.SrvUsers.GetUserByEmail(m.Email)
	if err != nil {
		res.Code = 15
		res.Type = "Error"
		res.Msg = "No se pudo obtener el correo electronico"
		logger.Error.Println("No se pudo obtener el correo electronico: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	if usr == nil {
		res.Code = 32
		res.Type = "Alert"
		res.Msg = "el email no pertenece a ninguna cuenta registrada"
		logger.Error.Println("el email no pertenece a ninguna cuenta registrada: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	usr.Password = ""
	res.Data = usr
	res.Code = 29
	res.Type = "success"
	res.Msg = "El correo si pertenece a una cuenta registrada"
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
