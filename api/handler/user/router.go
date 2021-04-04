package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UserRouter(app *fiber.App, db *sqlx.DB, tx string) {

	usr := Handler{DB: db, TxID: tx}
	api := app.Group("/api")
	v1 := api.Group("/v1/user")
	v1.Post("/register", usr.CreateUser)
/*	v1.Get("/user-exist", usr.ExistUser)
	v1.Get("/email-exist", usr.ExistEmail)*/
}