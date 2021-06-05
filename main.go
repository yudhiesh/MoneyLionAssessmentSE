package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/yudhiesh/api/controller"
	"github.com/yudhiesh/api/middleware"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.ResponseHeaders)
	router.Use(middleware.LogRequest)
	err := godotenv.Load(".env")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	dsn := db_user + ":" + db_password + "@/" + db_name
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
	router.HandleFunc("/feature", app.GetCanAccess).Methods("GET")
	router.HandleFunc("/feature", app.InsertFeature).Methods("POST")
	http.Handle("/", router)
	port := os.Getenv("PORT")
	portAddr := ":" + port
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	server := &http.Server{
		Handler: handler,
		Addr:    portAddr,
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
