package main

//go run cmd/server/main.go

import (
	"log"

	"github.com/Kingobhaisahb/nalini-art-gallery/config"
	"github.com/Kingobhaisahb/nalini-art-gallery/database"
	"github.com/Kingobhaisahb/nalini-art-gallery/routes"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	_, err := config.GetCloudinary()

	if err != nil {
		log.Fatal("Cloudinary configuration failed:", err)
	}

	log.Println("Cloudinary configured successfully")

	database.ConnectDatabase()

	router := gin.Default()

		router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
	}))

	routes.AuthRoutes(router)
	routes.PaintingRoutes(router)
	routes.PaintingImageRoutes(router)
	routes.PaintingVideoRoutes(router)
	routes.CartRoutes(router)
	routes.OrderRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Nalini Art Gallery Backend Running 🚀",
		})
	})

	log.Println("Server started on port 8080")

	router.Run(":8080")
}


