package main

import (
	"log"
	"os"

	// "rest-api/handlers"

	"rest-api/config"
	bioskop "rest-api/handlers/bioskop"
	home "rest-api/handlers/home"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	err error
)

func main() {
	// Membaca file .env
	if err = godotenv.Load(); err != nil {
		log.Fatal("Gagal membaca file .env: ", err)
	}
	if err := config.InitDatabase(); err != nil {
		log.Fatal(err)
	}
	defer config.DB.Close()

	// Membuat router Gin
	router := gin.Default()

	// API
	api := router.Group("/api")

	api.GET("/", home.GetHomeHandler)

	// API Bioskop
	api.GET("/bioskop", bioskop.GetBioskopHandler)
	api.POST("/bioskop", bioskop.CreateBioskopHandler)
	api.GET("/bioskop/:id", bioskop.GetDetailBioskopHandler)
	api.PUT("/bioskop/:id", bioskop.UpdateBioskopHandler)
	api.DELETE("/bioskop/:id", bioskop.DeleteBioskopHandler)

	// Menjalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Gagal menjalankan router: ", err)
	}
}
func init() {

}
