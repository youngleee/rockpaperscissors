package services

import (
	"database/sql"
	"fmt"
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

	var count int
	err := u.db.QueryRow(checkQuery, username).Scan(&count) // check if user exists
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists: %v", err)
	}
	if count > 0 {
		return nil, fmt.Errorf("user '%s' already exists", username)
	}

	insertQuery := `
	INSERT INTO users (username, total_coins, current_streak, games_played, games_won, created_at, updated_at)
	VALUES (?,0,0,0,0,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)
	`

	// create the user in database
	result, err := u.db.Exec(insertQuery, username)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get user ID: %v", err)
	}

	return &models.User{
		ID:            int(userID),
		Username:      username,
		TotalCoins:    0,
		CurrentStreak: 0,
		GamesPlayed:   0,
		GamesWon:      0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func (u *UserService) GetUser(username string) (*models.User, error) {
	query := `SELECT id, username, total_coins, current_streak, games_played, games_won, created_at, updated_at
	          FROM users
			  WHERE username = ?`

	// TODO: Execute query and scan results into user struct
	var user models.User

	err := u.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.TotalCoins,
		&user.CurrentStreak,
		&user.GamesPlayed,
		&user.GamesWon,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user '%s' not found", username)
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return &user, nil
}

func (u *UserService) UpdateUserStats(userID int, totalCoins int, currentStreak int, gamesPlayed int, gamesWon int) error {
	query := `UPDATE users
	          SET total_coins = ?, current_streak = ?, games_played = ?, games_won = ?, updated_at = CURRENT_TIMESTAMP
			  WHERE id = ?`
	result, err := u.db.Exec(query, totalCoins, currentStreak, gamesPlayed, gamesWon, userID)
	if err != nil {
		return fmt.Errorf("failed to update user stats: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check update result: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", userID)
	}

	return nil
}
