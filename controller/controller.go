package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/go-playground/validator"
	"github.com/yudhiesh/api/model"

	"github.com/yudhiesh/api/config"
)

func GetCanAccess(w http.ResponseWriter, r *http.Request) {
	var user model.UserCanAccess
	var responseFailure model.ResponseInfo
	var responseSuccess model.Response

	email := r.URL.Query().Get("email")
	featureName := r.URL.Query().Get("featureName")

	db := config.Connect()
	defer db.Close()
	stmt := `SELECT features.can_access FROM users INNER JOIN features ON users.id=features.user_id WHERE users.email=? AND features.feature_name=?`

	// Check if both url parameters have been set by the user
	if utf8.RuneCountInString(email) == 0 || utf8.RuneCountInString(featureName) == 0 {
		responseFailure.SetHeader(w, MissingURLParameter, http.StatusUnprocessableEntity)
		return
	}
	if err := db.QueryRow(stmt, email, featureName).Scan(&user.CanAccess); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Check if the sql query throws a ErrNoRows error
			responseFailure.SetHeader(w, NoMatchingRecordFound, http.StatusNotFound)
		} else {
			// Anything else return a StatusInternalServerError
			responseFailure.SetHeader(w, Error, http.StatusInternalServerError)
		}
		return
	} else {
		// If successfull return user object
		responseSuccess.SetHeader(w, Success, http.StatusOK, user)
	}
}

func InsertFeature(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var response model.ResponseInfo
	validate := validator.New()
	stmt := `UPDATE features INNER JOIN users ON features.user_id=users.id SET features.can_access=? WHERE users.email=? and features.feature_name=?`

	db := config.Connect()
	defer db.Close()

	// Decode body into user struct
	if err := parse(w, r, &user); err != nil {
		response.SetHeader(w, Error, http.StatusNotModified)
		return
	} else {
		// Validate struct to check if all fields are correct
		if err := validate.Struct(user); err != nil {
			response.SetHeader(w, MissingRequestParameter, http.StatusNotModified)
			return
		} else {
			// Check if can_access from the response and the database are different
			if _, access := CanAccessValue(db, *user.CanAccess, user.Email, user.FeatureName); !access {
				response.SetHeader(w, NoMatchingRecordFound, http.StatusNotModified)
				return

			} else if _, err = db.Exec(stmt, &user.CanAccess, &user.Email, &user.FeatureName); err != nil {
				// Update can_access in database
				response.SetHeader(w, Error, http.StatusNotModified)
				return
			} else {
				response.SetHeader(w, Success, http.StatusOK)
				return
			}

		}

	}

}
