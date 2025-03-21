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
	{
		URI:         "/users/{userId}/follow",
		Method:      "POST",
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/unfollow",
		Method:      "POST",
		Function:    controllers.UnfollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/followers",
		Method:      "GET",
		Function:    controllers.GetFollowers,
		RequireAuth: true,
	},
}
