package main

import (
	"fmt"
	"os"

	"github.com/PrinceNarteh/go-events-api/configs"
	"github.com/PrinceNarteh/go-events-api/controllers"
	"github.com/PrinceNarteh/go-events-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	fmt.Println(os.Getenv("MONGO_URI"))

	// initialize database
	configs.InitDB()

	server := gin.Default()
	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.Run(":4000")
}
