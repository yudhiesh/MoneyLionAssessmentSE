package main

import (
	"io/ioutil"
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
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	err := godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		errorLog.Fatal("$PORT is not set")
	}
	dsn := os.Getenv("DSN")
	if dsn == "" {
		errorLog.Fatal("$DATABASE_URL is not set")
	}
	db, err := controller.OpenDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	c, ioErr := ioutil.ReadFile("./schema.sql")
	sqlScript := string(c)
	if ioErr != nil {
		errorLog.Fatalf("Error loading SQL schema : %s", ioErr)
	}
	_, err = db.Exec(sqlScript)
	if err != nil {
		errorLog.Fatalf("Error executing SQL script : %s", err)
	}
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
