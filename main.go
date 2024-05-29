package main

import (
	"fmt"
	handler "hello/Handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")
	r := gin.Default()
	publicRoutes := r.Group("/public")
	//publicRoutes.Use(Middleware.CORSMiddleware())
	{
		publicRoutes.POST("/check", handler.HandleTelegramWebHook)
	}

	r.Run(":8989")

}
