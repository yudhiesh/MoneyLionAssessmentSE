package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&data)
	return err
}

// EmailExists takes in the users email and checks that their email exists
// before using it in any db query
func (a *Application) EmailExists(email string) (error, bool) {
	tx, err := a.DB.Begin()
	if err != nil {
		tx.Rollback()
		a.ErrorLog.Print(DBError)
		return err, false
	}
	row := tx.QueryRow("select email from users where email= ?", email)
	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return nil, true
	}
	err = tx.Commit()
	return err, false
}

// CanAccessValue looks the email up in the features and users tbale then checks
// if the canAccess passed in the request is different from the one in the
// database
func (a *Application) CanAccessValue(canAccess bool, email, featureName string) (error, bool) {
	tx, err := a.DB.Begin()
	if err != nil {
		tx.Rollback()
		a.ErrorLog.Print(DBError)
		return err, false
	}
	row := tx.QueryRow("SELECT features.can_access FROM features INNER JOIN users ON features.user_id=users.id WHERE users.email=? AND features.feature_name=?", email, featureName)
	rowResult := ""
	row.Scan(&rowResult)
	if rowResult == "0" {
		return nil, strconv.FormatBool(canAccess) != "false"
	} else if rowResult == "1" {
		return nil, strconv.FormatBool(canAccess) != "true"
	}
	err = tx.Commit()
	return err, false
}
