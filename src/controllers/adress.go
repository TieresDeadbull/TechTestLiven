package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddressController struct {
	DB   db.DBConnector
	Auth auth.Auth
}

// Função de cadastro de endereço
func (a *AddressController) CreateAddress(w http.ResponseWriter, r *http.Request) {

	var address models.Address

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

	db, err := a.DB.Connect()
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

	response.JSON(w, http.StatusCreated, address)
}

// Função de busca de endereço
func (a *AddressController) GetAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	addressID, err := strconv.ParseUint(params["addressId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return

	}
	db, err := a.DB.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	addressRepo := repositories.NewAddressesRepo(db)
	address, err := addressRepo.GetAddressesByID(addressID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, address)
}

// Função de listagem de todos os endereços
func (a *AddressController) ListAddresses(w http.ResponseWriter, r *http.Request) {
	db, err := a.DB.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	addressRepo := repositories.NewAddressesRepo(db)
	addresss, err := addressRepo.ListAddresses()

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, addresss)
}

// Função de atualização de endereço
func (a *AddressController) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var address models.Address

	params := mux.Vars(r)

	addressID, err := strconv.ParseUint(params["addressId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
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

	db, err := a.DB.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	addressRepo := repositories.NewAddressesRepo(db)
	if err = addressRepo.UpdateAddress(addressID, address); err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// Função de deleção de endereço
func (a *AddressController) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	addressID, err := strconv.ParseUint(params["addressId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return

	}

	db, err := a.DB.Connect()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	addressRepo := repositories.NewAddressesRepo(db)
	err = addressRepo.DeleteAddress(addressID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
