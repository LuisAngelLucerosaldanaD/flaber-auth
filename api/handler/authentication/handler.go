package authentication

import (
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"flaber-auth/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Handler struct {
	DB *sqlx.DB
	TxID string
}

func (h *Handler) Login(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	m := User{}
	err := c.BodyParser(&m)
	if err != nil {
		res.Code = 1
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	m.RealIP = c.IP()
	srvNxm := auth.NewServerAuth(h.DB, nil, h.TxID)
	token, code, err := srvNxm.SrvLogin.Login(m.Email, m.Cellphone, m.Password, m.RealIP)
	if err != nil {
		res.Code = 81
		res.Type = "Error"
		res.Msg = "usuario o contraseña incorrecta"
		logger.Error.Println("usuario o contraseña incorrectos: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = token
	res.Code = code
	res.Type = "success"
	res.Msg = "Procesado correctamente"
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}