package router

import (
	"golang-jwttoken/controller"
	"golang-jwttoken/middleware"
	"golang-jwttoken/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UserRepository,
	authenticationController *controller.AuthenticationController,
	usersController *controller.UserController,
	productController *controller.ProductController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	{
		authenticationRouter := router.Group("/authentication")
		{
			authenticationRouter.POST("/register", authenticationController.Register)
			authenticationRouter.POST("/login", authenticationController.Login)
		}
		usersRouter := router.Group("/users")
		{
			usersRouter.GET("", middleware.DeserializeUser(userRepository), usersController.GetUsers)
		}
		productRouter := router.Group("/product")
		{
			productRouter.GET("", productController.GetAllProduct)
			// productRouter.GET("/:productId", productController.GetProductById)
			productRouter.GET("/:productName", productController.GetProductByProductName)
			productRouter.POST("", productController.CreateNewProduct)
		}
	}
	return service
}
