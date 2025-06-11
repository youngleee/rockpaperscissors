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

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// setupTestDB creates a temporary test database
func setupTestDB(t *testing.T) *sql.DB {
	// Create temporary directory for test database
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, "test.db")

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

// setupTestRouter creates a test router with handlers
func setupTestRouter(db *sql.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	userHandler := NewUserHandler(db)

	// Setup routes
	api := router.Group("/api")
	api.POST("/users", userHandler.CreateUser)
	api.GET("/users/:username", userHandler.GetUser)
	api.GET("/stats/:username", userHandler.GetUserStats)
	api.GET("/leaderboard", userHandler.GetLeaderboard)

	return router
}

func TestUserHandler_CreateUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	router := setupTestRouter(db)

	t.Run("Success - Valid user creation", func(t *testing.T) {
		reqBody := models.CreateUserRequest{
			Username: "testuser",
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Check status code
		if w.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
		}

		// Parse response
		var response models.UserResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		// Verify response data
		if response.Username != "testuser" {
			t.Errorf("Expected username 'testuser', got '%s'", response.Username)
		}
		if response.TotalCoins != 0 {
			t.Errorf("Expected 0 coins for new user, got %d", response.TotalCoins)
		}
		if response.GamesPlayed != 0 {
			t.Errorf("Expected 0 games played for new user, got %d", response.GamesPlayed)
		}
	})

	t.Run("Error - Duplicate username", func(t *testing.T) {
		// First, create a user
		reqBody := models.CreateUserRequest{Username: "duplicate"}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Try to create the same user again
		req2 := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
		req2.Header.Set("Content-Type", "application/json")
		w2 := httptest.NewRecorder()
		router.ServeHTTP(w2, req2)

		// Check status code
		if w2.Code != http.StatusConflict {
			t.Errorf("Expected status %d for duplicate user, got %d", http.StatusConflict, w2.Code)
		}
	})

	t.Run("Error - Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d for invalid JSON, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Error - Empty username", func(t *testing.T) {
		reqBody := models.CreateUserRequest{Username: ""}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d for empty username, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestUserHandler_GetUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	router := setupTestRouter(db)

	// Create a test user first
	userHandler := NewUserHandler(db)
	testUser, err := userHandler.userService.CreateUser("getuser_test")
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	t.Run("Success - Get existing user", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/getuser_test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response models.UserResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		if response.Username != testUser.Username {
			t.Errorf("Expected username '%s', got '%s'", testUser.Username, response.Username)
		}
		if response.ID != testUser.ID {
			t.Errorf("Expected ID %d, got %d", testUser.ID, response.ID)
		}
	})

	t.Run("Error - User not found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/nonexistent", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d for non-existent user, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestUserHandler_GetUserStats(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	router := setupTestRouter(db)

	// Create a test user
	userHandler := NewUserHandler(db)
	_, err := userHandler.userService.CreateUser("statsuser")
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	t.Run("Success - Get user stats", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/stats/statsuser", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response models.UserStats
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		if response.Username != "statsuser" {
			t.Errorf("Expected username 'statsuser', got '%s'", response.Username)
		}
		if response.Rank <= 0 {
			t.Errorf("Expected rank > 0, got %d", response.Rank)
		}
	})

	t.Run("Error - User not found for stats", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/stats/nonexistent", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d for non-existent user stats, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestUserHandler_GetLeaderboard(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	router := setupTestRouter(db)

	t.Run("Success - Get empty leaderboard", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/leaderboard", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		// Check if leaderboard key exists
		if _, exists := response["leaderboard"]; !exists {
			t.Error("Response should contain 'leaderboard' key")
		}
		if _, exists := response["total_users"]; !exists {
			t.Error("Response should contain 'total_users' key")
		}
	})

	t.Run("Success - Get leaderboard with users", func(t *testing.T) {
		// Create test users with different coin amounts
		userHandler := NewUserHandler(db)

		user1, _ := userHandler.userService.CreateUser("leader1")
		user2, _ := userHandler.userService.CreateUser("leader2")

		// Update their stats to have different coin amounts
		userHandler.userService.UpdateUserStats(user1.ID, 100, 2, 5, 4)
		userHandler.userService.UpdateUserStats(user2.ID, 50, 1, 3, 2)

		req := httptest.NewRequest("GET", "/api/leaderboard", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response: %v", err)
		}

		// Verify we have users in leaderboard
		leaderboard, ok := response["leaderboard"].([]interface{})
		if !ok {
			t.Error("Leaderboard should be an array")
		}

		if len(leaderboard) < 2 {
			t.Errorf("Expected at least 2 users in leaderboard, got %d", len(leaderboard))
		}
	})
}

// TestUserHandler_Integration tests the full flow
func TestUserHandler_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	router := setupTestRouter(db)

	t.Run("Full user lifecycle", func(t *testing.T) {
		username := "integration_user"

		// 1. Create user
		reqBody := models.CreateUserRequest{Username: username}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("Failed to create user: status %d", w.Code)
		}

		// 2. Get user
		req2 := httptest.NewRequest("GET", "/api/users/"+username, nil)
		w2 := httptest.NewRecorder()
		router.ServeHTTP(w2, req2)

		if w2.Code != http.StatusOK {
			t.Fatalf("Failed to get user: status %d", w2.Code)
		}

		// 3. Get user stats
		req3 := httptest.NewRequest("GET", "/api/stats/"+username, nil)
		w3 := httptest.NewRecorder()
		router.ServeHTTP(w3, req3)

		if w3.Code != http.StatusOK {
			t.Fatalf("Failed to get user stats: status %d", w3.Code)
		}

		// 4. Check leaderboard includes user
		req4 := httptest.NewRequest("GET", "/api/leaderboard", nil)
		w4 := httptest.NewRecorder()
		router.ServeHTTP(w4, req4)

		if w4.Code != http.StatusOK {
			t.Fatalf("Failed to get leaderboard: status %d", w4.Code)
		}

		t.Log("✅ Full user lifecycle test passed!")
	})
}
