package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&data)
	return err
}

func EmailExists(db *sql.DB, email string) bool {
	row := db.QueryRow("select email from users where email= ?", email)
	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return true
	}
	return false
}

func CanAccessValue(db *sql.DB, canAccess bool, email, featureName string) (error, bool) {
	row := db.QueryRow("SELECT features.can_access FROM features INNER JOIN users ON features.user_id=users.id WHERE users.email=? AND features.feature_name=?", email, featureName)
	rowResult := ""
	row.Scan(&rowResult)
	err1 := errors.New("Failed to check value")
	if rowResult == "0" {
		return nil, strconv.FormatBool(canAccess) != "false"
	} else if rowResult == "1" {
		return nil, strconv.FormatBool(canAccess) != "true"
	}
	return err1, false
}
