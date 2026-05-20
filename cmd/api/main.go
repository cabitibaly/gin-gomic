package main

import (
	"github.com/cbitbaly/config"
	"github.com/cbitbaly/internal/database"
	"github.com/cbitbaly/internal/handlers"
	"github.com/cbitbaly/internal/repositories"
	"github.com/cbitbaly/internal/routes"
	"github.com/cbitbaly/internal/services"
	"github.com/cbitbaly/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	utils.InitializeJWTSecret(cfg.JWTSecret, cfg.RefreshTokenSecret)

	if err := database.Connect(cfg); err != nil {
		panic(err)
	}

	if err := database.Migration(); err != nil {
		panic(err)
	}

	db := database.GetDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	authRepo := repositories.NewRefreshTokenRepository(db)
	authService := services.NewAuthService(authRepo, userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	router := gin.Default()

	routes.UserRoutes(
		router,
		userHandler,
		authHandler,
		authService,
	)

	routes.PostRoutes(
		router,
		postHandler,
		authService,
	)

	router.Run()
}
