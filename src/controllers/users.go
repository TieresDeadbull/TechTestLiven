package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	ListUsers(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	DB   db.DBConnector
	Auth auth.Auth
}

// Função de cadastro de usuário
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

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

	if err = user.Prepare("register", &security.Encrypted{}); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := u.DB.Connect()
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

// Função de busca de usuário
func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return

	}
	db, err := u.DB.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUsersRepo(db)
	user, err := userRepo.GetUserByID(userID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// Função de listagem de todos os usuários
func (u *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	db, err := u.DB.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUsersRepo(db)
	users, err := userRepo.ListAllUsers()

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

// Função de atualização de usuário
func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := u.Auth.ExtractUserID(r)
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

	if err = user.Prepare("update", &security.Encrypted{}); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := u.DB.Connect()
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
func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return

	}

	tokenUserID, err := u.Auth.ExtractUserID(r)
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

	db, err := u.DB.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUsersRepo(db)
	err = userRepo.DeleteUser(userID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
