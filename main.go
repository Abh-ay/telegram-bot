package main

import (
	"encoding/json"
	"fmt"
	core "hello/Core"
	models "hello/Models"
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

}
func main() {
	fmt.Println("Hello world")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("failed to read body: %v", err)
		}
		var result models.Result
		if err := json.Unmarshal(body, &result); err != nil {
			log.Printf("failed to unmarshal body: %v", err)
			return
		}
		_, err = core.SendMessages(result)
		if err != nil {
			return
		}

	}
	http.HandleFunc("/", helloHandler)
	if err := http.ListenAndServe(":8989", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// r := gin.Default()
	// publicRoutes := r.Group("/public")
	// //publicRoutes.Use(Middleware.CORSMiddleware())
	// {
	// 	publicRoutes.GET("/check", handler.GetUpdates)
	// }
	//r.Run(":8989")

}
