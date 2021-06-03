package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/go-playground/validator"
	"github.com/yudhiesh/api/model"

	"github.com/yudhiesh/api/config"
)

func emailExists(db *sql.DB, email string) bool {
	row := db.QueryRow("select email from users where email= ?", email)
	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return true
	}
	return false
}

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
		responseFailure.Status = http.StatusUnprocessableEntity
		responseFailure.Message = MissingURLParameter
		json.NewEncoder(w).Encode(responseFailure)
		return
	}
	if err := db.QueryRow(stmt, email, featureName).Scan(&user.CanAccess); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Check if the sql query throws a ErrNoRows error
			responseFailure.Status = http.StatusNotFound
			responseFailure.Message = NoMatchingRecordFound
			json.NewEncoder(w).Encode(responseFailure)
			return
		} else {
			// Anything else return a StatusInternalServerError
			responseFailure.Status = http.StatusInternalServerError
			responseFailure.Message = Error
			json.NewEncoder(w).Encode(responseFailure)
			return
		}
	} else {
		// If successfull return user object
		responseSuccess.ResponseInfo.Status = http.StatusOK
		responseSuccess.ResponseInfo.Message = Success
		responseSuccess.Data = user
		json.NewEncoder(w).Encode(responseSuccess)
	}
}

func InsertFeature(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var response model.ResponseInfo
	validate := validator.New()
	stmt := `INSERT INTO features (user_id, feature_name, can_access) SELECT id, ?, ? FROM users WHERE email=?`

	db := config.Connect()
	defer db.Close()

	// Decode body into user struct
	if err := parse(w, r, &user); err != nil {
		response.Message = Error
		response.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	} else {
		// Validate struct to check if all fields are correct
		if err := validate.Struct(user); err != nil {
			response.Message = MissingRequestParameter
			response.Status = http.StatusUnprocessableEntity
			json.NewEncoder(w).Encode(response)
			return
		} else {
			// Check if the user exists in the table
			if !emailExists(db, user.Email) {
				response.Message = UserNotFound
				response.Status = http.StatusNotFound
				json.NewEncoder(w).Encode(response)
				return

			} else if _, err = db.Exec(stmt, &user.FeatureName, &user.CanAccess, &user.Email); err != nil {
				// Execute insert statement to database
				response.Message = Error
				response.Status = http.StatusInternalServerError
				json.NewEncoder(w).Encode(response)
				return
			} else {
				response.Message = Success
				response.Status = http.StatusOK
				json.NewEncoder(w).Encode(response)
				return
			}

		}

	}

}
