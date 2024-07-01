package routes

import (
	"api/src/controllers"
	"net/http"
)

// Aqui estão as rotas referentes a gestão de endereços
var addressRoutes = []Route{
	//Rota para criação de endereços ainda nao cadastrado
	{
		URI:          "/address/{userId}",
		Method:       http.MethodPost,
		Function:     controllers.CreateAddress,
		AuthRequired: true,
	},
	//Rota para vizualização de endereços cadastrados
	{
		URI:          "/user/address/{addressId}",
		Method:       http.MethodGet,
		Function:     controllers.GetAddress,
		AuthRequired: true,
	},
	//Rota para vizualização de endereços cadastrados filtrados por queryParam
	{
		URI:          "/user/address",
		Method:       http.MethodGet,
		Function:     controllers.GetAddress,
		AuthRequired: true,
	},
	//Rota para atualização de dados
	{
		URI:          "/address/{userId}/update/{addressId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateAddress,
		AuthRequired: true,
	},
	//Rota para deleção de endereço
	{
		URI:          "/address/{userId}/delete/{addressId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteAddress,
		AuthRequired: true,
	},
}
