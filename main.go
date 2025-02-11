package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/roadleft/roadleft-msgc/controllers"
	"github.com/roadleft/roadleft-msgc/initializers"
	"github.com/roadleft/roadleft-msgc/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	// os.Exit(-1)
}
