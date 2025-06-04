package services

import (
	"database/sql"
	"rockpaperscissors/internal/models"
	"time"
)

// UserService handles user-related operations
type UserService struct {
	db *sql.DB
}

// NewUserService creates a new user service
func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) CreateUser(username string) (*models.User, error) {
	checkQuery := `SELECT COUNT(*) FROM users WHERE username = ?`

	insertQuery := `
	INSERT INTO users (username, total_coins, current_streak, games_played, games_won, created_at, updated_at)
	VALUES (?,0,0,0,0,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)
	`

	// TODO: Add database execution logic here

	return &models.User{
		Username:      username,
		TotalCoins:    0,
		CurrentStreak: 0,
		GamesPlayed:   0,
		GamesWon:      0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}
