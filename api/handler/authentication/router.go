package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func AuthenticationRouter(app *fiber.App, db *sqlx.DB, tx string) {

	ln := Handler{DB: db, TxID: tx}

	api := app.Group("/api")
/*	v1 := api.Group("/v1/auth")
	v1.Post("/forgot-password", ln.ForgotPassword)
	v1.Post("/change-password", ln.ChangePassword)
	v1.Post("/password-policy", ln.PasswordPolicy)
*/
	v2 := api.Group("/v2/auth")
	v2.Post("/login", ln.Login)
}