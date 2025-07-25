package main

import (
	"fmt"
	"log"

	"github.com/abhidhanve/universal-dashboard/services/db_access/configs"
	"github.com/abhidhanve/universal-dashboard/services/db_access/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize environment configurations
	configs.InitEnvConfigs()

	// Get port from environment (fallback to 9081 if not set)
	port := configs.Env.Port
	if port == "" {
		port = "9081"
	}

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000", "http://localhost:8081"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	routes.SetupRoutes(router)

	fmt.Printf("🚀 DB Access Service (MVC Architecture) starting on port %s\n", port)
	fmt.Println("📋 Available endpoints:")
	fmt.Println("   • Health: GET /ping")
	fmt.Println("   • Allocate DB: POST /allocate")
	fmt.Println("   • Collections: GET /collections/:db")
	fmt.Println("   • Schema Detection: GET /detect-schema/:db/:collection")
	fmt.Println("   • Documents: GET /entries/:db/:collection")
	fmt.Println("   • CRUD: POST|GET|PUT|DELETE /entry/:db/:collection[/:id]")
	fmt.Println("   • API v1: /api/v1/*")
	fmt.Println("   • Info: GET /")

	// Start the server (bind to all interfaces for Railway)
	serverAddr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("🚀 Starting server on %s\n", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
