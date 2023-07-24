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

	// REVIEW: maybe pass the user repository to the favorite congress person repository is the best way
	favoriteCongressPersonRepository repository.FavoriteCongressPersonRepository = repository.NewFavoriteCongressPersonRepository()
	favoriteCongressPersonService    service.FavoriteCongressPersonService       = service.NewFavoriteCongressPersonService(favoriteCongressPersonRepository)
	favoriteCongressPersonController controller.FavoriteCongressPersonController = controller.NewFavoriteCongressPersonController(favoriteCongressPersonService)
)

func setHeaders(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	ctx.Next()
}

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	userAPI := api.NewUserAPI(userController)
	favoriteCongressPersonApi := api.NewFavoriteCongressPersonApi(favoriteCongressPersonController)

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

			favoriteCongresspersons := user.Group("/:id/favorite-congresspersons")
			{
				favoriteCongresspersons.GET("/", favoriteCongressPersonApi.GetFavoriteCongressPersons)
				favoriteCongresspersons.POST("/", favoriteCongressPersonApi.AddFavoriteCongressPerson)
			}
		}
	}


	server.Run("localhost:5000")
}