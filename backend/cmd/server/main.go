package main

import (
	"log"

	"github.com/Kingobhaisahb/nalini-art-gallery/config"
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	database.ConnectDatabase()

	router := gin.Default()

	routes.AuthRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Nalini Art Gallery Backend Running 🚀",
		})
	})

	log.Println("Server started on port 8080")

	router.Run(":8080")
}