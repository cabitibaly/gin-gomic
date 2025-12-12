package routes

import (
	"github.com/cbitbaly/internal/handlers"
	"github.com/cbitbaly/internal/services"
	"github.com/cbitbaly/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(
	r *gin.Engine,
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
	authService *services.AuthService,
) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.RegisterUserHandler)
		auth.POST("/login", authHandler.LoginUserHandler)
		auth.POST("/refresh-token", authHandler.RefreshTokenHandler)
	}

	users := r.Group("/users")
	users.Use(middlewares.AuthMiddleware(authService))
	{
		users.GET("/", userHandler.GetAllUserHandler)
		users.GET("/me", userHandler.GetMyInfoHandler)
		users.GET("/:id", userHandler.GetUserByIDHandler)
		users.PATCH("/:id", userHandler.UpdateUserHandler)
		users.DELETE("/:id", userHandler.DeleteUserHandler)
	}
}
