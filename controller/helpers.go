package controller

import (
	"encoding/json"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&data)
	return err
}
