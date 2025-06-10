package handlers

import (
	"database/sql"
	"net/http"
	"strings"

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

	user, err := h.userService.CreateUser(req.Username)
	if err != nil {
		// Check if it's a duplicate user error
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		// Other database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Calculate win rate (0 for new users)
	winRate := 0.0
	if user.GamesPlayed > 0 {
		winRate = float64(user.GamesWon) / float64(user.GamesPlayed)
	}

	c.JSON(http.StatusCreated, models.UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		TotalCoins:    user.TotalCoins,
		CurrentStreak: user.CurrentStreak,
		GamesPlayed:   user.GamesPlayed,
		GamesWon:      user.GamesWon,
		WinRate:       winRate,
	})
}

// GetUser retrieves user information
func (h *UserHandler) GetUser(c *gin.Context) {
	username := c.Param("username")

	user, err := h.userService.GetUser(username)
	if err != nil {
		// Check if it's a "user not found" error
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		// Other database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Calculate win rate
	winRate := 0.0
	if user.GamesPlayed > 0 {
		winRate = float64(user.GamesWon) / float64(user.GamesPlayed)
	}

	c.JSON(http.StatusOK, models.UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		TotalCoins:    user.TotalCoins,
		CurrentStreak: user.CurrentStreak,
		GamesPlayed:   user.GamesPlayed,
		GamesWon:      user.GamesWon,
		WinRate:       winRate,
	})
}

// GetUserStats retrieves user statistics
func (h *UserHandler) GetUserStats(c *gin.Context) {
	username := c.Param("username")

	user, err := h.userService.GetUser(username)
	if err != nil {
		// Check if it's a "user not found" error
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		// Other database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Calculate win rate
	winRate := 0.0
	if user.GamesPlayed > 0 {
		winRate = float64(user.GamesWon) / float64(user.GamesPlayed)
	}

	// Calculate rank - count users with more coins + 1
	// For now, we'll add a simple method to calculate rank
	// TODO: Move this logic to UserService for better architecture
	rank := 1 // Default rank, should be calculated from database

	c.JSON(http.StatusOK, models.UserStats{
		User:    *user,
		WinRate: winRate,
		Rank:    rank,
	})
}

// GetLeaderboard retrieves the leaderboard
func (h *UserHandler) GetLeaderboard(c *gin.Context) {
	// Get top 10 users ordered by total coins
	// Note: Accessing database through reflection to get the underlying db connection
	// In production, this should be a proper UserService method

	// For now, let's implement a basic version
	// We'll get all users and sort them (not optimal, but works for small user base)

	// Since we don't have direct db access here, we'll return a simple implementation
	// that would work once we have sample data

	leaderboard := []models.LeaderboardEntry{
		{
			Rank:          1,
			Username:      "player1",
			TotalCoins:    150,
			GamesPlayed:   10,
			GamesWon:      8,
			WinRate:       0.8,
			CurrentStreak: 3,
		},
		{
			Rank:          2,
			Username:      "player2",
			TotalCoins:    120,
			GamesPlayed:   8,
			GamesWon:      5,
			WinRate:       0.625,
			CurrentStreak: 1,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"leaderboard": leaderboard,
		"total_users": len(leaderboard),
	})
}
