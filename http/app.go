package http

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	http.Server
	host string
	port int
}

func NewApp(host string, port int, router *mux.Router) *App {
	return &App{
		host:   host,
		port:   port,
		Server: http.Server{Addr: ":" + strconv.Itoa(port), Handler: router},
	}
}

func (a *App) Run() error {
	return a.ListenAndServe()
}
