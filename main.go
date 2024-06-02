package main

import (
	"fmt"
	dbConnection "hello/DBConnection"
	handler "hello/Handler"
	utils "hello/Utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const filePath = "quries.sql"
const schemaFilePath = "schema.sql"

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

}

// func ListenSever() *http.Server {
// 	fmt.Println("Hello world")
// r := gin.Default()
// webHookRoutes := r.Group("")
// {
// 	webHookRoutes.POST("/", handler.WebHookHandler)
// }
// srv := &http.Server{
// 	Addr:    ":8989",
// 	Handler: r,
// }
// if err := srv.ListenAndServe(); err != nil {
// 	log.Printf("listen: %s\n", err)
// }
// return srv
// }

func main() {
	db := dbConnection.ConnectDB()
	//utils.InstallSchema(schemaFilePath, db)
	qMap := utils.ReadQueries(filePath)
	qr := utils.PrepareQueries(qMap, db)
	fmt.Println()
	handler.C.SetQueries(qr)
	r := gin.Default()
	webHookRoutes := r.Group("")
	{
		webHookRoutes.POST("/", handler.WebHookHandler)
	}
	srv := &http.Server{
		Addr:    ":8989",
		Handler: r,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}
