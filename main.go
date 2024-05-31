package main

import (
	"fmt"
	handler "hello/Handler"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
func main() {
	fmt.Println("Hello world")
	r := gin.Default()
	publicRoutes := r.Group("/public")
	//publicRoutes.Use(Middleware.CORSMiddleware())
	{
		publicRoutes.GET("/check", handler.HandleTelegramWebHook)
	}

	r.Run(":8989")

}
