package models

import "time"

// User represents a player in the rock-paper-scissors game
type User struct {
	ID           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	TotalCoins   int       `json:"total_coins" db:"total_coins"`
	CurrentStreak int      `json:"current_streak" db:"current_streak"`
	GamesPlayed  int       `json:"games_played" db:"games_played"`
	GamesWon     int       `json:"games_won" db:"games_won"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// UserStats represents calculated user statistics
type UserStats struct {
	User
	WinRate float64 `json:"win_rate"`
	Rank    int     `json:"rank"`
}

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
}

// UserResponse represents the public user information
type UserResponse struct {
	ID           int     `json:"id"`
	Username     string  `json:"username"`
	TotalCoins   int     `json:"total_coins"`
	CurrentStreak int     `json:"current_streak"`
	GamesPlayed  int     `json:"games_played"`
	GamesWon     int     `json:"games_won"`
	WinRate      float64 `json:"win_rate"`
} 