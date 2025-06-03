package models

import "time"

// Choice represents the possible choices in rock-paper-scissors
type Choice string

const (
	Rock     Choice = "rock"
	Paper    Choice = "paper"
	Scissors Choice = "scissors"
)

// GameResult represents the outcome of a game
type GameResult string

const (
	Win  GameResult = "win"
	Lose GameResult = "lose"
	Tie  GameResult = "tie"
)

// Game represents a single game round
type Game struct {
	ID               int        `json:"id" db:"id"`
	UserID           int        `json:"user_id" db:"user_id"`
	PlayerChoice     Choice     `json:"player_choice" db:"player_choice"`
	ComputerChoice   Choice     `json:"computer_choice" db:"computer_choice"`
	Result           GameResult `json:"result" db:"result"`
	CoinsEarned      int        `json:"coins_earned" db:"coins_earned"`
	StreakMultiplier int        `json:"streak_multiplier" db:"streak_multiplier"`
	PlayedAt         time.Time  `json:"played_at" db:"played_at"`
}

// PlayGameRequest represents the request to play a game
type PlayGameRequest struct {
	Username     string `json:"username" binding:"required"`
	PlayerChoice Choice `json:"player_choice" binding:"required"`
}

// PlayGameResponse represents the response after playing a game
type PlayGameResponse struct {
	PlayerChoice     Choice     `json:"player_choice"`
	ComputerChoice   Choice     `json:"computer_choice"`
	Result           GameResult `json:"result"`
	CoinsEarned      int        `json:"coins_earned"`
	StreakMultiplier int        `json:"streak_multiplier"`
	NewStreak        int        `json:"new_streak"`
	TotalCoins       int        `json:"total_coins"`
	Message          string     `json:"message"`
}

// LeaderboardEntry represents a player's position on the leaderboard
type LeaderboardEntry struct {
	Rank         int     `json:"rank"`
	Username     string  `json:"username"`
	TotalCoins   int     `json:"total_coins"`
	GamesPlayed  int     `json:"games_played"`
	GamesWon     int     `json:"games_won"`
	WinRate      float64 `json:"win_rate"`
	CurrentStreak int    `json:"current_streak"`
}

// IsValidChoice checks if the choice is valid
func (c Choice) IsValid() bool {
	return c == Rock || c == Paper || c == Scissors
}

// Beats returns true if this choice beats the other choice
func (c Choice) Beats(other Choice) bool {
	switch c {
	case Rock:
		return other == Scissors
	case Paper:
		return other == Rock
	case Scissors:
		return other == Paper
	default:
		return false
	}
} 