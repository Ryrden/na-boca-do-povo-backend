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

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	userAPI := api.NewUserAPI(userController)

	apiRoutes := server.Group("/api")
	{
		user := apiRoutes.Group("/user")
		{
			user.GET("/", userAPI.GetUsers)
			user.GET("/:id", userAPI.GetUser)
			user.POST("/", userAPI.Create)
			user.PUT("/:id", userAPI.Update)
			user.DELETE("/:id", userAPI.Delete)
		}
	}

	server.Run("localhost:5000")
}