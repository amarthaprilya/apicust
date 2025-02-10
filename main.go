// main.go
package main

import (
	"customer/db"
	"customer/handler"
	"customer/repository"
	"customer/service"
	"log"
	"os"

	_ "customer/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sweager Service API
// @description Sweager service API in Go using Gin framework
// @host localhost:8080/
func main() {
	database := db.ConnectDB()
	repo := &repository.CustomerRepository{DB: database}
	svc := &service.CustomerService{Repo: repo}
	handler := &handler.CustomerHandler{Service: svc}

	r := gin.Default()
	r.GET("/customers", handler.GetAll)
	r.GET("/customers/:id", handler.GetByID)
	r.POST("/customers", handler.Create)
	r.PUT("/customers/:id", handler.Update)
	r.DELETE("/customers/:id", handler.Delete)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin , Accept , X-Requested-With , Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization"},
		AllowMethods:    []string{"POST, OPTIONS, GET, PUT, DELETE"},
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Jalankan server pada port yang sudah ditentukan
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
