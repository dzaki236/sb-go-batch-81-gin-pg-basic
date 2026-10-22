package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "rest-api/handlers"
	handlers "rest-api/handlers/bioskop"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "DB_HOST"
	port     = "DB_PORT"
	user     = "DB_USER"
	password = "DB_PASSWORD"
	dbname   = "DB_NAME"
	sslmode  = "DB_SSLMODE"
)

var (
	db  *sql.DB
	err error
)

func main() {
	// Membaca file .env
	if err = godotenv.Load(); err != nil {
		log.Fatal("Gagal membaca file .env: ", err)
	}

	// Mengambil konfigurasi database
	dbHost := os.Getenv(host)
	dbPort := os.Getenv(port)
	dbUser := os.Getenv(user)
	dbPassword := os.Getenv(password)
	dbName := os.Getenv(dbname)
	dbSSLMode := os.Getenv(sslmode)

	// Memastikan konfigurasi tersedia
	if dbHost == "" || dbPort == "" ||
		dbUser == "" || dbPassword == "" ||
		dbName == "" || dbSSLMode == "" {
		log.Fatal("Konfigurasi database di .env belum lengkap")
	}

	// Membuat DSN PostgreSQL
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
		dbSSLMode,
	)

	// Membuka koneksi database
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Gagal membuka database: ", err)
	}
	defer db.Close()

	// Memastikan koneksi database aktif
	if err = db.Ping(); err != nil {
		log.Fatal("Gagal terhubung ke PostgreSQL: ", err)
	}

	log.Println("Successfully connected to database")

	// Membuat router Gin
	router := gin.Default()

	// API
	api := router.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Home berjalan",
		})
	})

	// API Bioskop
	api.POST("/bioskop", handlers.CreateBioskopHandler(db))

	// Menjalankan server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan router: ", err)
	}
}
