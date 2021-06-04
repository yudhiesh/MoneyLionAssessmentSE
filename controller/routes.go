package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/yudhiesh/api/middleware"
)

func (app *Application) routes() http.Handler {
	standardMiddleware := alice.New(middleware.LogRequest, middleware.ResponseHeaders)
	router := mux.NewRouter()
	router.HandleFunc("/feature", app.GetCanAccess).Methods("GET")
	router.HandleFunc("/feature", app.InsertFeature).Methods("POST")

	return standardMiddleware.Then(router)
}
