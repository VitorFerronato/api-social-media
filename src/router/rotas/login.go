package routes

import (
	"api/src/controllers"
)

var loginRoute = Route{
	URI:         "/login",
	Method:      "POST",
	Function:    controllers.Login,
	RequireAuth: false,
}
