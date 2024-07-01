package routes

import (
	"api/src/controllers"
	"net/http"
)

var addressController = controllers.AddressController{}

// Aqui estão as rotas referentes a gestão de endereços
var addressRoutes = []Route{
	//Rota para criação de endereços ainda nao cadastrado
	{
		URI:          "/address",
		Method:       http.MethodPost,
		Function:     addressController.CreateAddress,
		AuthRequired: false,
	},
	//Rota para vizualização de dados cadastrais
	{
		URI:          "/address/view/{addressId}",
		Method:       http.MethodGet,
		Function:     addressController.GetAddress,
		AuthRequired: false,
	},
	//Rota para atualização de dados
	{
		URI:          "/address/update/{addressId}",
		Method:       http.MethodPut,
		Function:     addressController.UpdateAddress,
		AuthRequired: false,
	},
	//Rota para deleção de endereço
	{
		URI:          "/address/delete/{addressId}",
		Method:       http.MethodDelete,
		Function:     addressController.DeleteAddress,
		AuthRequired: false,
	},
	//Rota para vizualização de dados cadastrais todos endereços cadastrados
	{
		URI:          "/addresses/view",
		Method:       http.MethodGet,
		Function:     addressController.ListAddresses,
		AuthRequired: false,
	},
}
