package routes

import (
	"github.com/cbitbaly/internal/handlers"
	"github.com/cbitbaly/internal/services"
	"github.com/cbitbaly/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func PostRoutes(
	r *gin.Engine,
	postHandler *handlers.PostHandler,
	authService *services.AuthService,
) {
	posts := r.Group("/posts")
	posts.Use(middlewares.AuthMiddleware(authService))
	{
		posts.POST("/", postHandler.CreatePostHandler)
		posts.GET("/", postHandler.GetAllPostsHandler)
		posts.GET("/:id", postHandler.GetPostByIdHandler)
		posts.GET("/my-posts", postHandler.GetAllPostsByUserIDHandler)
		posts.PATCH("/:id", postHandler.UpdatePostHandler)
		posts.DELETE("/:id", postHandler.DeletePostHandler)
	}
}
