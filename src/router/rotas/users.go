package routes

import (
	"api/src/controllers"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      "POST",
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      "GET",
		Function:    controllers.GetUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      "GET",
		Function:    controllers.GetUserById,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      "PUT",
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      "DELETE",
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
