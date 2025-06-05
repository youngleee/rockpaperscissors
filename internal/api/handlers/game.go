package handlers

import (
	"database/sql"
	"net/http"

	"rockpaperscissors/internal/models"
	"rockpaperscissors/internal/services"

	"github.com/gin-gonic/gin"
)

// GameHandler handles game-related requests
type GameHandler struct {
	gameService *services.GameService
}

// NewGameHandler creates a new game handler
func NewGameHandler(db *sql.DB) *GameHandler {
	return &GameHandler{
		gameService: services.NewGameService(db),
	}
}

// PlayGame handles playing a rock-paper-scissors game
func (h *GameHandler) PlayGame(c *gin.Context) {
	// TODO: Implement game logic
	c.JSON(http.StatusOK, gin.H{
		"message": "PlayGame endpoint - to be implemented",
	})
}

// GetUserGames retrieves game history for a user
func (h *GameHandler) GetUserGames(c *gin.Context) {
	// TODO: Implement game history retrieval
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"message":  "GetUserGames endpoint - to be implemented",
		"username": username,
	})
}

func (h *GameHandler) PlayGame(c *gin.Context) {
	// Step 1: parse and validate the request
	var req models.PlayGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 2: Validate player choice
	if !req.PlayerChoice.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid choice, must be 'rock', 'paper', or 'scissors'"})
		return
	}

	// Step 3: Play the game using the game service
	response, err := h.gameService.PlayGame(req.Username, req.PlayerChoice)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Step 4: Return the Game Result
	c.JSON(http.StatusOK, response)
}
