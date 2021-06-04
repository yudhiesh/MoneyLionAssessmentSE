package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/rs/cors"
	"github.com/yudhiesh/api/controller"
	"github.com/yudhiesh/api/middleware"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.ResponseHeaders)
	router.Use(middleware.LogRequest)
	dsn := flag.String("dsn", "root:password@/assessment", "MySQL data source name")
	flag.Parse()
	db, err := openDB(*dsn)
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
	router.HandleFunc("/feature", app.GetCanAccess).Methods("GET")
	router.HandleFunc("/feature", app.InsertFeature).Methods("POST")
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
	infoLog.Printf("Connected to port %s", port)
	errorLog.Fatal(server.ListenAndServe())
}

// Returns a sql.DB connection pool for a given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Connections are established lazily as and when needed for the first time
	// db.Ping creates a connection and we check that there isn't any errors
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
