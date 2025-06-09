package main

import (
	"cp_tracker/routes"

	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 設置Gin模式
	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	// 使用 wire 初始化應用
	app, err := routes.InitializeApp()
	if err != nil {
		log.Fatalf("Cannot initialize app: %v", err)
	}

	// 設置所有 Router
	r := routes.InitializeRoute(&app)

	// 啟動服務
	r.Run(":8080")
}
