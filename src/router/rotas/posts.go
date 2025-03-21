package routes

import (
	"api/src/controllers"
)

var postRoutes = []Route{
	{
		URI:         "/posts",
		Method:      "POST",
		Function:    controllers.CreatePost,
		RequireAuth: true,
	},
	{
		URI:         "/posts",
		Method:      "GET",
		Function:    controllers.GetPosts,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{postId}",
		Method:      "GET",
		Function:    controllers.GetPostById,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{postId}",
		Method:      "PUT",
		Function:    controllers.UpdatePost,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{postId}",
		Method:      "DELETE",
		Function:    controllers.DeletePost,
		RequireAuth: true,
	},
}
