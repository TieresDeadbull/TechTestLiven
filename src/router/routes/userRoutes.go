package routes

import (
	"api/src/controllers"
	"net/http"
)

var userController = controllers.UserController{}

// Aqui estão as rotas referentes a gestão de usuários
var userRoutes = []Route{
	//Rota para criação de usuário ainda nao cadastrado
	{
		URI:          "/user",
		Method:       http.MethodPost,
		Function:     userController.CreateUser,
		AuthRequired: false,
	},
	//Rota para vizualização de dados cadastrais
	{
		URI:          "/user/view/{userId}",
		Method:       http.MethodGet,
		Function:     userController.GetUser,
		AuthRequired: true,
	},
	//Rota para atualização de dados
	{
		URI:          "/user/update/{userId}",
		Method:       http.MethodPut,
		Function:     userController.UpdateUser,
		AuthRequired: true,
	},
	//Rota para deleção de conta
	{
		URI:          "/user/delete/{userId}",
		Method:       http.MethodDelete,
		Function:     userController.DeleteUser,
		AuthRequired: true,
	},
	//Rota para vizualização de dados cadastrais todos usuarios cadastrados
	{
		URI:          "/users/view",
		Method:       http.MethodGet,
		Function:     userController.ListUsers,
		AuthRequired: true,
	},
}
