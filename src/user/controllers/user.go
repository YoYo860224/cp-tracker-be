package controllers

import (
	"cp_tracker/user/models"
	"cp_tracker/user/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: *userService,
	}
}

/*
CreateUser creates a new user.
*/
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	// Bind the JSON input to the User model
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := c.userService.CreateUser(&user); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(201, user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	user, err := c.userService.GetUserByUid(uid)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(200, user)
}
