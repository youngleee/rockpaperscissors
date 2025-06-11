package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"rockpaperscissors/internal/database"
	"rockpaperscissors/internal/models"
	"rockpaperscissors/internal/services"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// setupGameTestDB creates a temporary test database for game tests
func setupGameTestDB(t *testing.T) *sql.DB {
	// Create temporary directory for test database
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, "game_test.db")

	// Open test database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		t.Fatalf("Failed to enable foreign keys: %v", err)
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	return db
}

// setupGameTestRouter creates a test router with game handlers
func setupGameTestRouter(db *sql.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	gameHandler := NewGameHandler(db)

	// Setup routes
	api := router.Group("/api")
	api.POST("/play", gameHandler.PlayGame)
	api.GET("/users/:username/games", gameHandler.GetUserGames)

	return router
}

func TestGameHandler_PlayGame(t *testing.T) {
	db := setupGameTestDB(t)
	defer db.Close()

	router := setupGameTestRouter(db)

	// Create a test user first
	userService := services.NewUserService(db)
	_, err := userService.CreateUser("gamer123")
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	t.Run("Success - Play rock and win", func(t *testing.T) {
		reqBody := models.PlayGameRequest{
			Username:     "gamer123",
			PlayerChoice: models.Rock,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/play", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Check status code
		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		// Parse response
		var response models.PlayGameResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		// Verify response structure
		if response.PlayerChoice != models.Rock {
			t.Errorf("Expected player choice 'rock', got '%s'", response.PlayerChoice)
		}
		if !response.ComputerChoice.IsValid() {
			t.Errorf("Computer choice should be valid, got '%s'", response.ComputerChoice)
		}
		if response.Result != models.Win && response.Result != models.Lose && response.Result != models.Tie {
			t.Errorf("Invalid game result: %s", response.Result)
		}
		if response.Message == "" {
			t.Error("Response should contain a message")
		}
		if response.TotalCoins < 0 {
			t.Errorf("Total coins should not be negative, got %d", response.TotalCoins)
		}

		t.Logf("Game result: Player=%s, Computer=%s, Result=%s, Coins=%d",
			response.PlayerChoice, response.ComputerChoice, response.Result, response.CoinsEarned)
	})

	t.Run("Success - Play paper", func(t *testing.T) {
		reqBody := models.PlayGameRequest{
			Username:     "gamer123",
			PlayerChoice: models.Paper,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/play", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response models.PlayGameResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		if response.PlayerChoice != models.Paper {
			t.Errorf("Expected player choice 'paper', got '%s'", response.PlayerChoice)
		}
	})

	t.Run("Success - Play scissors", func(t *testing.T) {
		reqBody := models.PlayGameRequest{
			Username:     "gamer123",
			PlayerChoice: models.Scissors,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/play", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response models.PlayGameResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		if response.PlayerChoice != models.Scissors {
			t.Errorf("Expected player choice 'scissors', got '%s'", response.PlayerChoice)
		}
	})

	t.Run("Error - User not found", func(t *testing.T) {
		reqBody := models.PlayGameRequest{
			Username:     "nonexistentuser",
			PlayerChoice: models.Rock,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/play", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d for non-existent user, got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("Error - Invalid choice", func(t *testing.T) {
		// Use raw JSON to send invalid choice
		invalidJson := `{"username": "gamer123", "player_choice": "dynamite"}`

		req := httptest.NewRequest("POST", "/api/play", bytes.NewBufferString(invalidJson))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d for invalid choice, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Error - Empty username", func(t *testing.T) {
		reqBody := models.PlayGameRequest{
			Username:     "",
			PlayerChoice: models.Rock,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/play", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d for empty username, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Error - Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/play", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d for invalid JSON, got %d", http.StatusBadRequest, w.Code)
		}
	})

	// Test that user stats are updated after playing
	t.Run("Success - User stats updated after game", func(t *testing.T) {
		// Get initial user stats
		initialUser, _ := userService.GetUser("gamer123")
		initialGamesPlayed := initialUser.GamesPlayed
		initialTotalCoins := initialUser.TotalCoins

		// Play a game
		reqBody := models.PlayGameRequest{
			Username:     "gamer123",
			PlayerChoice: models.Rock,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/play", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Game request failed: status %d", w.Code)
		}

		// Get updated user stats
		updatedUser, err := userService.GetUser("gamer123")
		if err != nil {
			t.Fatalf("Failed to get updated user: %v", err)
		}

		// Verify stats were updated
		if updatedUser.GamesPlayed != initialGamesPlayed+1 {
			t.Errorf("Games played should increase by 1: initial=%d, updated=%d",
				initialGamesPlayed, updatedUser.GamesPlayed)
		}

		// Total coins should be >= initial (could stay same on loss, increase on win)
		if updatedUser.TotalCoins < initialTotalCoins {
			t.Errorf("Total coins should not decrease: initial=%d, updated=%d",
				initialTotalCoins, updatedUser.TotalCoins)
		}

		t.Logf("Stats updated: Games %d→%d, Coins %d→%d",
			initialGamesPlayed, updatedUser.GamesPlayed,
			initialTotalCoins, updatedUser.TotalCoins)
	})
}

func TestGameHandler_GetUserGames(t *testing.T) {
	db := setupGameTestDB(t)
	defer db.Close()

	router := setupGameTestRouter(db)

	// Create a test user
	userService := services.NewUserService(db)
	_, err := userService.CreateUser("gamehistoryuser")
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	t.Run("Success - Get empty game history", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/gamehistoryuser/games", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Should return 200 OK with empty games array
		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		// Verify response contains expected keys
		if _, exists := response["username"]; !exists {
			t.Error("Response should contain 'username' key")
		}
		if _, exists := response["games"]; !exists {
			t.Error("Response should contain 'games' key")
		}
		if _, exists := response["total_games"]; !exists {
			t.Error("Response should contain 'total_games' key")
		}

		username, ok := response["username"].(string)
		if !ok || username != "gamehistoryuser" {
			t.Errorf("Expected username 'gamehistoryuser', got '%v'", response["username"])
		}

		games, ok := response["games"].([]interface{})
		if !ok {
			// Handle case where empty slice might be null in JSON
			if response["games"] == nil {
				// This is expected for empty game history
				t.Logf("Empty game history returned as null (expected)")
			} else {
				t.Errorf("Games should be an array or null, got %T", response["games"])
			}
		} else {
			if len(games) != 0 {
				t.Errorf("Expected 0 games for new user, got %d", len(games))
			}
		}
	})

	t.Run("Success - Get game history with games", func(t *testing.T) {
		// Create a user and play some games to generate history
		userService := services.NewUserService(db)
		_, err := userService.CreateUser("activegamer")
		if err != nil {
			t.Fatalf("Failed to create test user: %v", err)
		}

		// Play a few games
		gameChoices := []models.Choice{models.Rock, models.Paper}
		for _, choice := range gameChoices {
			reqBody := models.PlayGameRequest{
				Username:     "activegamer",
				PlayerChoice: choice,
			}
			jsonBody, _ := json.Marshal(reqBody)

			req := httptest.NewRequest("POST", "/api/play", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Failed to play game: status %d", w.Code)
			}
		}

		// Now get the game history
		req := httptest.NewRequest("GET", "/api/users/activegamer/games", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		games, ok := response["games"].([]interface{})
		if !ok {
			t.Error("Games should be an array")
		}

		if len(games) != 2 {
			t.Errorf("Expected 2 games in history, got %d", len(games))
		}

		totalGames, ok := response["total_games"].(float64)
		if !ok || int(totalGames) != 2 {
			t.Errorf("Expected total_games to be 2, got %v", response["total_games"])
		}

		t.Logf("Game history retrieved: %d games", len(games))
	})

	t.Run("Error - User not found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/nonexistentuser/games", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d for non-existent user, got %d", http.StatusNotFound, w.Code)
		}
	})
}

// TestGameHandler_Integration tests the full game flow
func TestGameHandler_Integration(t *testing.T) {
	db := setupGameTestDB(t)
	defer db.Close()

	router := setupGameTestRouter(db)

	t.Run("Multiple games with streak building", func(t *testing.T) {
		// Create user
		userService := services.NewUserService(db)
		_, err := userService.CreateUser("streakmaster")
		if err != nil {
			t.Fatalf("Failed to create test user: %v", err)
		}

		// Play multiple games to test streak mechanics
		choices := []models.Choice{models.Rock, models.Paper, models.Scissors}

		for i, choice := range choices {
			reqBody := models.PlayGameRequest{
				Username:     "streakmaster",
				PlayerChoice: choice,
			}
			jsonBody, _ := json.Marshal(reqBody)

			req := httptest.NewRequest("POST", "/api/play", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Game %d failed: status %d", i+1, w.Code)
			}

			var response models.PlayGameResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("Failed to parse game %d response: %v", i+1, err)
			}

			t.Logf("Game %d: %s vs %s = %s (Streak: %d, Coins: +%d)",
				i+1, choice, response.ComputerChoice, response.Result,
				response.NewStreak, response.CoinsEarned)
		}

		// Verify final user state
		finalUser, err := userService.GetUser("streakmaster")
		if err != nil {
			t.Fatalf("Failed to get final user state: %v", err)
		}

		if finalUser.GamesPlayed != 3 {
			t.Errorf("Expected 3 games played, got %d", finalUser.GamesPlayed)
		}

		t.Logf("Final stats: Games=%d, Wins=%d, Coins=%d, Streak=%d",
			finalUser.GamesPlayed, finalUser.GamesWon,
			finalUser.TotalCoins, finalUser.CurrentStreak)
	})
}

// TestGameLogic_Consistency tests that game outcomes are deterministic given the same inputs
func TestGameLogic_Consistency(t *testing.T) {
	db := setupGameTestDB(t)
	defer db.Close()

	userService := services.NewUserService(db)

	// Create test user
	_, err := userService.CreateUser("consistency_test")
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	t.Run("Rock Paper Scissors logic verification", func(t *testing.T) {
		// Test all possible combinations to ensure logic is correct
		testCases := []struct {
			player   models.Choice
			computer models.Choice
			expected models.GameResult
		}{
			{models.Rock, models.Scissors, models.Win},
			{models.Rock, models.Paper, models.Lose},
			{models.Rock, models.Rock, models.Tie},
			{models.Paper, models.Rock, models.Win},
			{models.Paper, models.Scissors, models.Lose},
			{models.Paper, models.Paper, models.Tie},
			{models.Scissors, models.Paper, models.Win},
			{models.Scissors, models.Rock, models.Lose},
			{models.Scissors, models.Scissors, models.Tie},
		}

		gameLogic := services.NewGameLogicService()

		for _, tc := range testCases {
			result := gameLogic.DetermineWinner(tc.player, tc.computer)
			if result != tc.expected {
				t.Errorf("Wrong result for %s vs %s: expected %s, got %s",
					tc.player, tc.computer, tc.expected, result)
			}
		}
	})

	t.Run("Coin calculation logic", func(t *testing.T) {
		gameLogic := services.NewGameLogicService()

		// Test base coin earning (10 coins for win)
		baseCoins := gameLogic.CalculateCoinsEarned(models.Win, 0)
		if baseCoins != 10 {
			t.Errorf("Expected 10 base coins for win, got %d", baseCoins)
		}

		// Test no coins for loss
		loseCoins := gameLogic.CalculateCoinsEarned(models.Lose, 0)
		if loseCoins != 0 {
			t.Errorf("Expected 0 coins for loss, got %d", loseCoins)
		}

		// Test no coins for tie
		tieCoins := gameLogic.CalculateCoinsEarned(models.Tie, 0)
		if tieCoins != 0 {
			t.Errorf("Expected 0 coins for tie, got %d", tieCoins)
		}

		// Test streak multiplier
		streakCoins := gameLogic.CalculateCoinsEarned(models.Win, 2) // Streak 2 = 3x multiplier
		if streakCoins != 30 {
			t.Errorf("Expected 30 coins with streak 2 (3x multiplier), got %d", streakCoins)
		}
	})
}
