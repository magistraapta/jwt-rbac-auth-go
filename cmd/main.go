package main

import (
	"log"
	"rbac-go/db"
	"rbac-go/internal/auth"
	"rbac-go/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	router := gin.Default()

	auth.SetupAuthRoutes(router, db)
	product.SetupProductRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
