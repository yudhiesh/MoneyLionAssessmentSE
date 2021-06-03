package model

import (
	"encoding/json"
	"net/http"
)

type UserCanAccess struct {
	CanAccess bool `json:"can_access"`
}

type User struct {
	FeatureName string `json:"featureName" validate:"required"`
	Email       string `json:"email" validate:"required"`
	CanAccess   *bool  `json:"can_access" validate:"required"`
}

type ResponseInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	ResponseInfo ResponseInfo  `json:"response"`
	Data         UserCanAccess `json:"data"`
}

func (r *ResponseInfo) SetHeader(w http.ResponseWriter, message string, status int) {
	// Set header with status, message and encode it into the response header
	r.Status = status
	r.Message = message
	json.NewEncoder(w).Encode(r)
}

func (r *Response) SetHeader(w http.ResponseWriter, message string, status int, data UserCanAccess) {
	// Set header with status, message and encode it into the response header
	r.ResponseInfo.Status = status
	r.ResponseInfo.Message = message
	r.Data = data
	json.NewEncoder(w).Encode(r)
}
