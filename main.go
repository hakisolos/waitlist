package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hakisolos/waitlist/config"
	"github.com/hakisolos/waitlist/controllers"
	_ "github.com/hakisolos/waitlist/docs"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Waitlist API
// @version 1.0
// @description API documentation for the Waitlist service
// @host localhost:3001
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config.ConnDB()
	app := gin.Default()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/", controllers.TestController)
	app.POST("/join", controllers.JoinController)
	app.GET("/users", controllers.GetUsersController)

	app.Run(":3001")
	fmt.Println(os.Getenv("MONGO_URI"))
}
