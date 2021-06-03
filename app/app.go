package app

import (
	"github.com/gorilla/mux"
	"github.com/yudhiesh/api/controller"
	"github.com/yudhiesh/api/database"
)

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/feature", controller.GetCanAccess).Methods("GET")
	a.Router.HandleFunc("/feature", controller.InsertFeature).Methods("POST")
}
