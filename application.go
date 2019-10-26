package webframework

import (
	"fmt"
	"net/http"
)

type application struct {
	server   http.Server
	settings Settings
}

func NewApplication(settings Settings) application {
	app := application{settings: settings}

	app.server = http.Server{
		Addr:    settings.Addr,
		Handler: app,
	}

	return app
}

func (app application) Run() {
	fmt.Printf("Listening on %s\n", app.server.Addr)
	app.server.ListenAndServe()
}
