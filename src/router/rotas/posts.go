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
	{
		URI:         "/users/{userId}/posts",
		Method:      "GET",
		Function:    controllers.GetPostsByUser,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{postId}/like",
		Method:      "POST",
		Function:    controllers.LikeInPost,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{postId}/unlike",
		Method:      "POST",
		Function:    controllers.UnlikeInPost,
		RequireAuth: true,
	},
}
