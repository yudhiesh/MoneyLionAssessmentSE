package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response of the GetCanAccess query
type UserCanAccess struct {
	CanAccess bool `json:"can_access"`
}

// User information that is sent in the InsertFeature request
type User struct {
	FeatureName string `json:"featureName" validate:"required"`
	Email       string `json:"email" validate:"required"`
	CanAccess   *bool  `json:"can_access" validate:"required"`
	// store pointer of the bool as default value of bool is false even with no value set
}

// Plain response
type Response struct{}

// Set the header of the successfull request with a specified format which is "data"
func (u *UserCanAccess) SetHeader(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// Set the error for the responses with the body returnded
func (u *UserCanAccess) SetError(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		u.SetHeader(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	u.SetHeader(w, http.StatusBadRequest, nil)
}

// Set the header for the response with no body returned
func (r *Response) SetHeader(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
