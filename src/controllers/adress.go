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

// Função de cadastro de endereço
func CreateAddress(w http.ResponseWriter, r *http.Request) {

	var address models.Address

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
		response.Err(w, http.StatusForbidden,
			errors.New("não permitida criação de endereço para outro usuário"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(body, &address); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = address.Prepare(); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	addressRepo := repositories.NewAddressesRepo(db)
	address.ID, err = addressRepo.Create(address)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	newUserAddress := repositories.NewUserAddressRepo(db, *repositories.NewUsersRepo(db), *repositories.NewAddressesRepo(db))
	if err := newUserAddress.CreateUserAddresses(userID, address.ID); err != nil {
		return
	}
	response.JSON(w, http.StatusCreated, address)
}

// Função de busca de endereço
func GetAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	queries := r.URL.Query()

	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		response.Err(w, http.StatusUnauthorized, err)
		return
	}

	addressId, ok := params["addressId"]
	if ok && addressId != "" {
		addressID, err := strconv.ParseUint(addressId, 10, 64)
		if err != nil {
			response.Err(w, http.StatusBadRequest, err)
			return
		}
		db, err := db.Connect()
		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Close()

		addressRepo := repositories.NewAddressesRepo(db)
		//Tratado internamente para só buscar endereços do usuarios por meio do ID extraido do token
		address, err := addressRepo.GetAddressesByID(tokenUserID, addressID)

		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}

		response.JSON(w, http.StatusOK, address)

	} else {
		db, err := db.Connect()
		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Close()

		addressRepo := repositories.NewAddressesRepo(db)
		address, err := addressRepo.GetAddressesByFilter(tokenUserID, queries)

		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}

		response.JSON(w, http.StatusOK, address)
	}
}

// Função de atualização de endereço
func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var address models.Address

	params := mux.Vars(r)

	addressID, err := strconv.ParseUint(params["addressId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}
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
		response.Err(w, http.StatusForbidden,
			errors.New("não permitida edição de outro endereço que não seja o seu"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(body, &address); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	addressRepo := repositories.NewAddressesRepo(db)
	if err = addressRepo.UpdateAddress(userID, addressID, address); err != nil {
		//internamente, por meio da query, caso o endereço nao tenho sido criado
		//pelo usuario nenhuma atualização é efetuada
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Função de deleção de endereço
// onde só é possivel remover o endereço associado ao seu usuario
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	addressID, err := strconv.ParseUint(params["addressId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return

	}

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
			errors.New("não permitida deleção de outro endereço que não seja o seu"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userAddressRepo := repositories.NewUserAddressRepo(db, *repositories.NewUsersRepo(db), *repositories.NewAddressesRepo(db))
	err = userAddressRepo.DeleteAddress(userID, addressID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
