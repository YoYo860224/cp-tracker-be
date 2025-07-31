package routes

import (
	"cp_tracker/user/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// App struct holds the dependencies for the application
func InitializeRoute(app *App) *gin.Engine {
	// Create a new Gin router
	r := gin.Default()

	// 設定 CORS，允許所有 localhost 來源
	corsConfig := cors.Config{
		AllowOrigins: []string{
			"http://localhost:4200",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	// Set trusted proxies to nil for local development
	r.SetTrustedProxies(nil)

	// API routes
	api := r.Group("/api/v1")
	user := api.Group("/user")
	userData := api.Group("/user-data")

	// User routes
	user.POST("", app.UserController.Register)
	user.POST("/login", app.UserController.Login)
	user.GET("", middlewares.JWTAuthMiddleware(), app.UserController.GetUser)
	user.PUT("", middlewares.JWTAuthMiddleware(), app.UserController.UpdateUserInfo)
	user.PUT("/password", middlewares.JWTAuthMiddleware(), app.UserController.UpdatePassword)

	// User data routes
	userData.GET("", middlewares.JWTAuthMiddleware(), app.UserDataController.GetUserData)
	userData.PUT("", middlewares.JWTAuthMiddleware(), app.UserDataController.UpdateUserData)

	return r
}
