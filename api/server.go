package api

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	version = "0.0.2"
	website = "https://luislucero.dev - https://juancx.dev"
	banner  = `
███████╗██╗░░░░░░█████╗ ██████╗ ███████╗██████╗░░░░░░░░█████╗░██╗░░░██╗████████╗██╗░░██╗
██╔════╝██║░░░░░██╔══██╗██╔══██╗██╔════╝██╔══██╗░░░░░░██╔══██╗██║░░░██║╚══██╔══╝██║░░██║
█████╗░░██║░░░░░███████║██████╔╝█████╗░░██████╔╝█████╗███████║██║░░░██║░░░██║░░░███████║
██╔══╝░░██║░░░░░██╔══██║██╔══██╗██╔══╝░░██╔══██╗╚════╝██╔══██║██║░░░██║░░░██║░░░██╔══██║
██║░░░░░███████╗██║░░██║██████╔╝███████╗██║░░██║░░░░░░██║░░██║╚██████╔╝░░░██║░░░██║░░██║
╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═════╝░╚══════╝╚═╝░░╚═╝░░░░░░╚═╝░░╚═╝░╚═════╝░░░░╚═╝░░░╚═╝░░╚═╝
                                                                                        `

	description = `Flaber Auth - Port: %s
by Flaber S.A.C
Version: %s
%s`
)

type server struct {
	listening string
	app       string
	fb        *fiber.App
}

func newServer(listening int, app string, fb *fiber.App) *server {
	return &server{fmt.Sprintf(":%d", listening), app, fb}
}

func (srv *server) Start() {
	color.Blue(banner)
	color.Cyan(fmt.Sprintf(description, srv.listening, version, website))
	log.Fatal(srv.fb.Listen(srv.listening))
}
