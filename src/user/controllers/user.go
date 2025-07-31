package controllers

import (
	"cp_tracker/user/models"
	"cp_tracker/user/services"
	"cp_tracker/user/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService       services.UserService
	inviteCodeService services.InviteCodeService
}

func NewUserController(
	userService *services.UserService,
	inviteCodeService *services.InviteCodeService) *UserController {
	return &UserController{
		userService:       *userService,
		inviteCodeService: *inviteCodeService,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	// 定義請求結構體
	var req struct {
		models.User
		InviteCode string `json:"inviteCode"`
	}

	// Bind the JSON input to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// 檢查邀請碼
	if err := c.inviteCodeService.ValidateInviteCode(req.InviteCode, req.Email); err != nil {
		ctx.JSON(403, gin.H{"error": "Invalid invite code or email not allowed"})
		return
	}

	// 創建用戶
	if err := c.userService.CreateUser(&req.User); err != nil {
		// 有可能是 email 重複
		if err.Error() == "duplicate email" {
			ctx.JSON(400, gin.H{"error": "User with this email already exists"})
			return
		}

		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	// 產生 JWT
	token, err := utils.GenerateJWT(req.User.Uid.Hex(), req.User.Email)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	userDto := req.User.ToDto()
	userDto.Token = token
	ctx.JSON(201, userDto)
}

func (c *UserController) Login(ctx *gin.Context) {
	// 定義請求結構體
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the JSON input to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// 根據 email 獲取用戶
	user, err := c.userService.GetUserByEmail(req.Email)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}
	if !c.userService.CheckPassword(user, req.Password) {
		ctx.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// 產生 JWT
	token, err := utils.GenerateJWT(user.Uid.Hex(), user.Email)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	userDto := user.ToDto()
	userDto.Token = token
	ctx.JSON(200, userDto)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	// 從上下文中獲取當前用戶的 uid
	loginUid := ctx.GetString("uid")

	// 查詢當前用戶
	user, err := c.userService.GetUserByUid(loginUid)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, user.ToDto())
}

func (c *UserController) UpdateUserInfo(ctx *gin.Context) {
	// 從上下文中獲取當前用戶的 uid
	loginUid := ctx.GetString("uid")

	// 定義請求結構體
	var req struct {
		DisplayName string `json:"displayName"`
	}

	// Bind the JSON input to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// 根據 uid 獲取用戶
	user, err := c.userService.GetUserByUid(loginUid)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// 更新顯示名稱
	user.DisplayName = req.DisplayName
	if err := c.userService.UpdateUser(user); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update display name"})
		return
	}

	ctx.JSON(200, user.ToDto())
}

func (c *UserController) UpdatePassword(ctx *gin.Context) {
	// 從上下文中獲取當前用戶的 uid
	loginUid := ctx.GetString("uid")

	// 定義請求結構體
	var req struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	// Bind the JSON input to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// 根據 uid 獲取用戶
	user, err := c.userService.GetUserByUid(loginUid)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// 檢查舊密碼是否正確
	if !c.userService.CheckPassword(user, req.OldPassword) {
		ctx.JSON(401, gin.H{"error": "Old password is incorrect"})
		return
	}

	// 更新密碼
	user.Password = req.NewPassword
	if err := c.userService.UpdatePassword(user.Uid, req.NewPassword); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update password"})
		return
	}

	ctx.JSON(200, user.ToDto())
}
