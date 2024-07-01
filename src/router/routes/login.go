package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginController = controllers.LoginController{}

var loginRoute = Route{
	//Rota de login usuário
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     loginController.LoginUser,
	AuthRequired: false,
}
