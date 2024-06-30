package routes

import (
	"api/src/controllers"
	"net/http"
)

// Aqui estão as rotas referentes a gestão de usuários
var userRoutes = []Route{
	//Rota para criação de usuário ainda nao cadastrado
	{
		URI:          "/user",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	//Rota para vizualização de dados cadastrais
	{
		URI:          "/user/view/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		AuthRequired: true,
	},
	//Rota para atualização de dados
	{
		URI:          "/user/update/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: true,
	},
	//Rota para deleção de conta
	{
		URI:          "/user/delete/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: true,
	},
	//Rota para vizualização de dados cadastrais todos usuarios cadastrados
	{
		URI:          "/users/view",
		Method:       http.MethodGet,
		Function:     controllers.ListUsers,
		AuthRequired: true,
	},
}
