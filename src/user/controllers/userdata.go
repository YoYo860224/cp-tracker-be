package controllers

import (
	"cp_tracker/user/models"
	"cp_tracker/user/services"

	"github.com/gin-gonic/gin"
)

type UserDataController struct {
	userDataService *services.UserDataService
}

func NewUserDataController(userDataService *services.UserDataService) *UserDataController {
	return &UserDataController{userDataService: userDataService}
}

func (c *UserDataController) GetUserData(ctx *gin.Context) {
	uid := ctx.GetString("uid")

	userData, err := c.userDataService.GetUserData(uid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get user data"})
		return
	}
	ctx.JSON(200, userData)
}

func (c *UserDataController) UpdateUserData(ctx *gin.Context) {
	uid := ctx.GetString("uid")

	var items map[string]interface{}
	if err := ctx.ShouldBindJSON(&items); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	userData := models.UserData{
		Uid:   uid,
		Items: items,
	}
	if err := c.userDataService.UpdateUserData(&userData); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update user data"})
		return
	}
	ctx.JSON(200, userData)
}
