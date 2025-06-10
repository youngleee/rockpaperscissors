package handlers

import (
	"database/sql"
	"net/http"

	"rockpaperscissors/internal/models"
	"rockpaperscissors/internal/services"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related requests
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(db),
	}
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Return 400 Bad Request with the validation error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Continue implementation - call user service and handle response
}

// GetUser retrieves user information
func (h *UserHandler) GetUser(c *gin.Context) {
	// TODO: Implement user retrieval
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"message":  "GetUser endpoint - to be implemented",
		"username": username,
	})
}

// GetUserStats retrieves user statistics
func (h *UserHandler) GetUserStats(c *gin.Context) {
	// TODO: Implement user stats retrieval
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"message":  "GetUserStats endpoint - to be implemented",
		"username": username,
	})
}

// GetLeaderboard retrieves the leaderboard
func (h *UserHandler) GetLeaderboard(c *gin.Context) {
	// TODO: Implement leaderboard
	c.JSON(http.StatusOK, gin.H{
		"message": "GetLeaderboard endpoint - to be implemented",
	})
}
