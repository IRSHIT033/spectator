package main

import (
	"time"

	"github.com/gin-gonic/gin"
	_authHandler "spectator.main/auth/transport/http"
	_authUsecase "spectator.main/auth/usecase"
	_configRepo "spectator.main/config/repository/mongo_repository"
	_configHandler "spectator.main/config/transport/http"
	_configUsecase "spectator.main/config/usecase"
	"spectator.main/internals/bootstrap"
	_userRepo "spectator.main/user/repository/mongo_repository"
	_userHandler "spectator.main/user/transport/http"
	_userUsecase "spectator.main/user/usecase"
)

func main() {

	app := bootstrap.App()

	config := app.Config

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	timeoutContext := time.Duration(config.ContextTimeout) * time.Second

	database := app.Mongo.Database(config.DBname)

	ginRouter := router.Group("api/v1")

	userRepo := _userRepo.NewMongoRepository(database)
	userUseCase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	_userHandler.NewUserHandler(ginRouter, userUseCase)
	authUseCase := _authUsecase.NewAuthUsecase(userRepo, timeoutContext)
	_authHandler.NewAuthHandler(config, ginRouter, authUseCase)

	configRepo := _configRepo.NewMongoRepository(database)
	configUseCase := _configUsecase.NewConfigUsecase(configRepo, userRepo, timeoutContext)
	_configHandler.NewConfigHandler(config, ginRouter, configUseCase)

	router.Run(":8080")
}
