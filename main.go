package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryrden/na-boca-do-povo-backend/api"
	"github.com/ryrden/na-boca-do-povo-backend/controller"
	"github.com/ryrden/na-boca-do-povo-backend/service"
	"github.com/ryrden/na-boca-do-povo-backend/repository"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func setHeaders(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	ctx.Next()
}

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	userAPI := api.NewUserAPI(userController)

	apiRoutes := server.Group("/api")
	apiRoutes.Use(setHeaders)
	{
		user := apiRoutes.Group("/user")
		{
			user.GET("/", userAPI.GetUsers)
			user.GET("/:id", userAPI.GetUser)
			user.POST("/", userAPI.Create)
			user.PUT("/:id", userAPI.Update)
			user.DELETE("/:id", userAPI.Delete)
			user.POST("/favorite-congressperson/:id", userAPI.AddFavoriteCongressPerson)
		}
	}


	server.Run("localhost:5000")
}