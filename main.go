package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zanuardinovanda/go-testing/config"
	"github.com/zanuardinovanda/go-testing/controllers"
	"github.com/zanuardinovanda/go-testing/middleware"
	"github.com/zanuardinovanda/go-testing/repository"
	"github.com/zanuardinovanda/go-testing/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	jwtService     service.JWTService         = service.NewJWTService()
	userService    service.UserService        = service.NewUserService(userRepository)
	authService    service.AuthService        = service.NewAuthService(userRepository)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController = controllers.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	r.Run()
}
