package user

import (
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"flaber-auth/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Handler struct {
	DB *sqlx.DB
	TxID string
}
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	res := models.Response{Error: false}
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
	usr, cod, err := srvNxm.SrvUsers.CreateUser(m.ID, m.Name, m.LastName, m.Password, m.EmailNotifications, m.IdentificationNumber, m.Cellphone)
	if err != nil {
		res.Code = cod
		res.Type = "Error"
		res.Msg = "don't create user"
		logger.Error.Println("No se pudo crear el usuario: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	usr.Password = ""
	res.Data = usr
	res.Code = cod
	res.Type = "success"
	res.Msg = "El usuario ha sido creado exitosamente"
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

// TODO implement method change password
func (h *Handler) ChangePassword(c *fiber.Ctx) error {
	res := models.Response{Error: false}
	m := requestChangePassword{}
	err := c.BodyParser(&m)
	if err != nil {
		res.Code = 1
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Code = 29
	res.Type = "success"
	res.Msg = "La contraseña se ha cambiado exitosamente"
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

// TODO implement method recovery password
func (h *Handler) RecoveryPassword(c *fiber.Ctx) error {
	res := models.Response{Error: false}
	m := requestRecoveryPassword{}
	err := c.BodyParser(&m)
	if err != nil {
		res.Code = 1
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Code = 29
	res.Type = "success"
	res.Msg = "Se ha recuperado correctamente la contraseña"
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) ExistEmail(c *fiber.Ctx) error {
	res := models.Response{Error: false}
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
