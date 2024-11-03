package controllers

import (
	"encoding/json"
	"net/http"
)

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, payload interface{}, status int) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, map[string]string{"message": message}, code)
}
