package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/yudhiesh/api/controller"
	"github.com/yudhiesh/api/middleware"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.ResponseHeaders)
	router.Use(middleware.LogRequest)
	err := godotenv.Load(".env")
	dsn := os.Getenv("DSN")
	port := os.Getenv("PORT")
	_ := os.Getenv("DB_HOST")
	db, err := controller.OpenDB(dsn)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	app := &controller.Application{
		DB:       db,
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}
	addr := ":" + port
	server := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	router.HandleFunc("/feature", app.GetCanAccess).Methods("GET")
	router.HandleFunc("/feature", app.InsertFeature).Methods("POST")
	http.Handle("/", router)
	infoLog.Printf("Connected to port %s", port)
	errorLog.Fatal(server.ListenAndServe())
}
