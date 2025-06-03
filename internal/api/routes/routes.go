package routes

import (
	"database/sql"

	"rockpaperscissors/internal/api/handlers"
	"rockpaperscissors/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the API routes
func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Initialize handlers
	gameHandler := handlers.NewGameHandler(db)
	userHandler := handlers.NewUserHandler(db)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Rock Paper Scissors API is running"})
	})

	// API routes group
	api := router.Group("/api")
	{
		// Apply common middleware to API routes
		api.Use(middleware.JSONMiddleware())
		api.Use(middleware.ErrorHandler())

		// User management
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:username", userHandler.GetUser)
		api.GET("/stats/:username", userHandler.GetUserStats)

		// Game endpoints
		api.POST("/play", gameHandler.PlayGame)
		api.GET("/leaderboard", userHandler.GetLeaderboard)
		
		// Game history (optional)
		api.GET("/users/:username/games", gameHandler.GetUserGames)
	}

	// Serve static files for web frontend (if needed)
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")
	
	// Web frontend route (optional)
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Rock Paper Scissors",
		})
	})
} 