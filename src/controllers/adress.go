package controllers

import (
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

// Função de cadastro de endereço
func CreateAddress(w http.ResponseWriter, r *http.Request) {

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

	response.JSON(w, http.StatusCreated, address)
}

// Função de busca de endereço
func GetAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	addressID, err := strconv.ParseUint(params["addressId"], 10, 64)
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
	address, err := addressRepo.GetAddressesByID(addressID)

	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, address)
}

// Função de listagem de todos os endereços
func ListAddresses(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
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
func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var address models.Address

	params := mux.Vars(r)

	addressID, err := strconv.ParseUint(params["addressId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	// tokenUserID, err := auth.ExtractUserID(r)
	// if err != nil {
	// 	response.Err(w, http.StatusUnauthorized, err)
	// 	return
	// }

	// if userID != tokenUserID {
	// 	fmt.Println(userID, tokenUserID)
	// 	response.Err(w, http.StatusForbidden,
	// 		errors.New("não permitida edição de outro endereço que não seja o seu"))
	// 	return
	// }

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(body, &address); err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	// if err = address.Prepare("update"); err != nil {
	// 	response.Err(w, http.StatusBadRequest, err)
	// 	return
	// }

	db, err := db.Connect()
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
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	addressID, err := strconv.ParseUint(params["addressId"], 10, 64)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return

	}

	// tokenAddressID, err := auth.ExtractAddressID(r)
	// if err != nil {
	// 	response.Err(w, http.StatusUnauthorized, err)
	// 	return
	// }

	// if addressID != tokenAddressID {
	// 	fmt.Println(addressID, tokenAddressID)
	// 	response.Err(w, http.StatusForbidden,
	// 		errors.New("não permitida deleção de outro endereço que não seja o seu"))
	// 	return
	// }

	db, err := db.Connect()
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
