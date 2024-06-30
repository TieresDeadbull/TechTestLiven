package routes

import (
	"api/src/controllers"
	"net/http"
)

// Aqui estão as rotas referentes a gestão de endereços
var addressRoutes = []Route{
	//Rota para criação de endereços ainda nao cadastrado
	{
		URI:          "/address",
		Method:       http.MethodPost,
		Function:     controllers.CreateAddress,
		AuthRequired: false,
	},
	//Rota para vizualização de dados cadastrais
	{
		URI:          "/address/view/{addressId}",
		Method:       http.MethodGet,
		Function:     controllers.GetAddress,
		AuthRequired: false,
	},
	//Rota para atualização de dados
	{
		URI:          "/address/update/{addressId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateAddress,
		AuthRequired: false,
	},
	//Rota para deleção de endereço
	{
		URI:          "/address/delete/{addressId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteAddress,
		AuthRequired: false,
	},
	//Rota para vizualização de dados cadastrais todos endereços cadastrados
	{
		URI:          "/addresses/view",
		Method:       http.MethodGet,
		Function:     controllers.ListAddresses,
		AuthRequired: false,
	},
}
