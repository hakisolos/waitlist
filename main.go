package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hakisolos/waitlist/config"
	"github.com/hakisolos/waitlist/controllers"
	"github.com/joho/godotenv"
)

// @title Waitlist API
// @version 1.0
// @description API documentation for the Waitlist service
// @host localhost:3001
// @BasePath /
func main() {
	_ = godotenv.Load()

	config.ConnDB()
	app := gin.Default()

	app.GET("/", controllers.TestController)
	app.POST("/join", controllers.JoinController)
	app.GET("/users", controllers.GetUsersController)

	app.Run(":3000")
	fmt.Println(os.Getenv("MONGO_URI"))
}
