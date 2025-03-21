package routes

import (
	"api/src/controllers"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      "POST",
		Function:    controllers.CreateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users",
		Method:      "GET",
		Function:    controllers.GetUsers,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      "GET",
		Function:    controllers.GetUserById,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      "PUT",
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      "DELETE",
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
}
