package routes

import (
	"github.com/gin-gonic/gin"
)

// App struct holds the dependencies for the application
func InitializeRoute(app *App) *gin.Engine {
	// Create a new Gin router
	r := gin.Default()

	// Set trusted proxies to nil for local development
	r.SetTrustedProxies(nil)

	// API routes
	api := r.Group("/api/v1")
	users := api.Group("/users")

	// User routes
	users.POST("", app.UserController.CreateUser)
	users.GET("/:uid", app.UserController.GetUser)

	return r
}
