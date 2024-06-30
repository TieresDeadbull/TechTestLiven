package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	//Rota de login usuário
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	AuthRequired: false,
}
