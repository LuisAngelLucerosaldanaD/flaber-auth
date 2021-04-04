package api

import (
	"flaber-auth/api/handler/authentication"
	"flaber-auth/api/handler/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

func routes(db *sqlx.DB, loggerHttp bool, allowedOrigins string) *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("fluber")
	prometheus.RegisterAt(app, "/metrics")

	app.Use(recover.New())
	app.Use(prometheus.Middleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST",
	}))
	if loggerHttp {
		app.Use(logger.New())
	}
	TxID := uuid.New().String()
	authentication.AuthenticationRouter(app, db, TxID)
	user.UserRouter(app, db, TxID)

	return app
}

