package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	http.Server
	host string
	port string
}

func NewApp(host, port string, router *mux.Router) *App {
	return &App{
		host:   host,
		port:   port,
		Server: http.Server{Addr: ":" + port, Handler: router},
	}
}

func (a *App) Run() error {
	return a.ListenAndServe()
}
