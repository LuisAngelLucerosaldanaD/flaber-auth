package api

import "flaber-auth/internal/dbx"

func Start(port int, app string, loggerHttp bool, allowedOrigins string) {

	db := dbx.GetConnection()
	defer db.Close()

	r := routes(db, loggerHttp, allowedOrigins)
	server := newServer(port, app, r)
	server.Start()
}
