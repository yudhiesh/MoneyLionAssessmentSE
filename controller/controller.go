package controller

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/go-playground/validator"
	"github.com/yudhiesh/api/model"
)

type Application struct {
	DB       *sql.DB
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

// GetCanAccess receives the email and featureName from the response and returns
// whether or not the user associated with the email has access to the feature
func (a *Application) GetCanAccess(w http.ResponseWriter, r *http.Request) {
	var user model.UserCanAccess
	var responseFailure model.ResponseInfo
	var responseSuccess model.Response

	email := r.URL.Query().Get("email")
	featureName := r.URL.Query().Get("featureName")

	tx, err := a.DB.Begin()
	if err != nil {
		tx.Rollback()
		a.ErrorLog.Print(DBError)
		return
	}
	stmt := `SELECT features.can_access FROM users INNER JOIN features ON users.id=features.user_id WHERE users.email=? AND features.feature_name=?`

	// Check if both url parameters have been set by the user
	if utf8.RuneCountInString(email) == 0 || utf8.RuneCountInString(featureName) == 0 {
		tx.Rollback()
		responseFailure.SetHeader(w, MissingURLParameter, http.StatusUnprocessableEntity)
		a.ErrorLog.Printf(MissingURLParameter)
		return
	}
	if err := tx.QueryRow(stmt, email, featureName).Scan(&user.CanAccess); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Check if the sql query throws a ErrNoRows error
			tx.Rollback()
			responseFailure.SetHeader(w, NoMatchingRecordFound, http.StatusNotFound)
			a.ErrorLog.Printf(NoMatchingRecordFound)
		} else {
			// Anything else return a StatusInternalServerError
			tx.Rollback()
			responseFailure.SetHeader(w, Error, http.StatusInternalServerError)
			a.ErrorLog.Printf(Error)
		}
		return
	} else {
		// If successfull return user object
		tx.Commit()
		responseSuccess.SetHeader(w, Success, http.StatusOK, user)
		a.InfoLog.Printf(Success)
	}
}

// InsertFeature receives featureName, email and enable from the response then
// switches the users access to a particular feature
func (a *Application) InsertFeature(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var response model.ResponseInfo
	validate := validator.New()
	stmt := `UPDATE features INNER JOIN users ON features.user_id=users.id SET features.can_access=? WHERE users.email=? and features.feature_name=?`

	tx, err := a.DB.Begin()
	if err != nil {
		tx.Rollback()
		a.ErrorLog.Printf(DBError)
		return
	}

	// Decode body into user struct
	if err := parse(w, r, &user); err != nil {
		tx.Rollback()
		response.SetHeader(w, Error, http.StatusNotModified)
		a.ErrorLog.Printf(Error)
		return
	} else {
		// Validate struct to check if all fields are correct
		if err := validate.Struct(user); err != nil {
			tx.Rollback()
			response.SetHeader(w, MissingRequestParameter, http.StatusNotModified)
			a.ErrorLog.Printf(MissingRequestParameter)
			return
		} else {
			// Check if can_access from the response and the database are different
			if _, access := a.CanAccessValue(*user.CanAccess, user.Email, user.FeatureName); !access {
				tx.Rollback()
				response.SetHeader(w, NoMatchingRecordFound, http.StatusNotModified)
				a.ErrorLog.Printf(NoMatchingRecordFound)
				return

			} else if _, err = a.DB.Exec(stmt, &user.CanAccess, &user.Email, &user.FeatureName); err != nil {
				// Check if updating can_access failed
				tx.Rollback()
				response.SetHeader(w, Error, http.StatusNotModified)
				a.ErrorLog.Printf(Error)
				return
			} else {
				tx.Commit()
				response.SetHeader(w, Success, http.StatusOK)
				a.InfoLog.Printf(Success)
				return
			}

		}

	}

}
