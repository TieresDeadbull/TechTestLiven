package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Função de cadastro de usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUsersRepo(db)
	user.ID, err = userRepo.Create(user)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login User"))
}

// Função de busca de usuário
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return

	}

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		fmt.Println(userID, tokenUserID)
		response.Err(w, http.StatusForbidden,
			errors.New("não permitida consulta de outro usuário que não seja o seu"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userAddressRepo := repositories.NewUserAddressRepo(db, *repositories.NewUsersRepo(db), *repositories.NewAddressesRepo(db))
	user, err := userAddressRepo.GetUserWithAddresses(userID)

	fmt.Println(user)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// Função de atualização de usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		fmt.Println(userID, tokenUserID)
		response.Err(w, http.StatusForbidden,
			errors.New("não permitida edição de outro usuário que não seja o seu"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(body, &user); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUsersRepo(db)
	if err = userRepo.UpdateUser(userID, user); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Função de deleção de usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return

	}

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		fmt.Println(userID, tokenUserID)
		response.Err(w, http.StatusForbidden,
			errors.New("não permitida deleção de outro usuário que não seja o seu"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userAddressRepo := repositories.NewUserAddressRepo(db, *repositories.NewUsersRepo(db), *repositories.NewAddressesRepo(db))
	err = userAddressRepo.DeleteUser(userID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
