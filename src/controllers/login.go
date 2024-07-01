package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LoginController struct {
	DB       db.DBConnector
	UserRepo repositories.UserRepository
	Security security.Security
	Auth     auth.Auth
}

// Função de login do usuário onde serão feitas as validações de senha, email, entre outras
func (lc *LoginController) LoginUser(w http.ResponseWriter, r *http.Request) {
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

	db, err := lc.DB.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userFromDB, err := lc.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = lc.Security.VerifyPass(userFromDB.Passphrase, user.Passphrase); err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := lc.Auth.CreateToken(userFromDB.ID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(token)
	w.Write([]byte("Você está logado."))
}
