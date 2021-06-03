package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/rs/cors"
	"github.com/yudhiesh/api/controller"
	"github.com/yudhiesh/api/middleware"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.ResponseMiddleware)
	router.HandleFunc("/feature", controller.GetCanAccess).Methods("GET")
	router.HandleFunc("/feature", controller.InsertFeature).Methods("POST")
	http.Handle("/", router)
	port := ":3000"
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	server := &http.Server{
		Handler: handler,
		Addr:    port,
		// keep-alives last a minute instead of 3 minutes
		IdleTimeout: time.Minute,
		// Short ReadTimeout prevents SLowloris attacks
		ReadTimeout: 5 * time.Second,
		// Prevent the data that the handler returns from taking too long to write
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Connected to port %s", port)
	log.Fatal(server.ListenAndServe())
}
