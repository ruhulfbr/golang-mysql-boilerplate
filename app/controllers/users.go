package controllers

import (
	"earn-expense/app/auth"
	"earn-expense/app/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	users, err := user.FindAllUsers(db)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, users, http.StatusOK)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user.Prepare()

	err = user.Validate("")

	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	createUser, err := user.CreateUser(db)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}

	RespondJSON(w, createUser, http.StatusOK)
}

func GetUserById(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	// string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{}
	existedUser, err := user.FindUserByID(db, uint32(userId))
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondJSON(w, existedUser, http.StatusOK)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	// string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	tokenID, err := auth.ExtractTokenID(r)

	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if tokenID != uint32(userId) {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	user.Prepare()

	err = user.Validate("")
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedUser, err := user.UpdateAUser(db, uint32(userId))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondJSON(w, updatedUser, http.StatusOK)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	// string to int
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if tokenID != 0 && tokenID != uint32(userId) {
		RespondError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	user := models.User{}
	_, err = user.DeleteAUser(db, uint32(userId))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondJSON(w, map[string]string{"message": "Delete."}, http.StatusOK)
}
