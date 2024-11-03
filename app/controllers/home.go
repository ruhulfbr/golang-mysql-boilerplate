package controllers

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

func Home(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	response := map[string]string{}
	response["status"] = "Success"
	response["message"] = "Welcome To This Awesome API 2."

	RespondJSON(w, response, http.StatusOK)
}
