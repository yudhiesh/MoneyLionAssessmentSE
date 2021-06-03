package config

import "database/sql"

func Connect() *sql.DB {
	DBDriver := "mysql"
	DBUser := "root"
	DBPass := "password"
	DBName := "assessment"
	db, err := sql.Open(DBDriver, DBUser+":"+DBPass+"@/"+DBName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
