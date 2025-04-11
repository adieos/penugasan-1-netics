package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	up_at := time.Now()

	// simulate env loading
	if err := loadEnv(); err != nil {
		fmt.Printf("Warning: %v\n", err.Error())
		fmt.Printf("Continuing without .env file (jangan tp cuek ws)\n")
	}

	fmt.Println("Setting up server...")
	server := gin.Default()
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "mau nyari apa to",
		})
	})

	envPentingJanganKeekspos := os.Getenv("ENV_PENTING")

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"nama":      "Athallah Rajendra Wibisono",
			"nrp":       "5025231170",
			"status":    "UP",
			"timestamp": time.Now().Format(time.RFC3339),
			"uptime":    time.Since(up_at).String(),
			"env":       fmt.Sprintf("ini env jangan sampai ke publik: %v", envPentingJanganKeekspos),
		})
	})

	// simulate static assets and start server
	server.Static("/assets", "./assets")
	port := os.Getenv("PORT") // dipake ae ws env ne awikwok
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = "0.0.0.0:" + port
	}

	fmt.Printf("Starting server on %s\n", serve)
	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func loadEnv() error {
	// get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	// get the directory of the executable
	ex, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error getting executable path: %v", err)
	}
	exPath := filepath.Dir(ex)

	// list of possible locations for .env file
	envLocations := []string{
		filepath.Join(cwd, ".env"),
		filepath.Join(exPath, ".env"),
		"/var/www/penugasan-1-netics/.env",
	}

	// try to load .env from each location
	for _, loc := range envLocations {
		err := godotenv.Load(loc)
		if err == nil {
			fmt.Printf("Loaded .env from: %s\n", loc)
			return nil
		}
	}

	return fmt.Errorf(".env gaada wok yg bener")
}
