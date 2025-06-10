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

// GetLeaderboard retrieves top users ordered by total coins
func (u *UserService) GetLeaderboard(limit int) ([]models.LeaderboardEntry, error) {
	if limit <= 0 {
		limit = 10 // Default to top 10
	}

	query := `SELECT username, total_coins, current_streak, games_played, games_won, created_at, updated_at
	          FROM users 
			  ORDER BY total_coins DESC, games_won DESC 
			  LIMIT ?`

	rows, err := u.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query leaderboard: %v", err)
	}
	defer rows.Close()

	var leaderboard []models.LeaderboardEntry
	rank := 1

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Username,
			&user.TotalCoins,
			&user.CurrentStreak,
			&user.GamesPlayed,
			&user.GamesWon,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan leaderboard row: %v", err)
		}

		// Calculate win rate
		winRate := 0.0
		if user.GamesPlayed > 0 {
			winRate = float64(user.GamesWon) / float64(user.GamesPlayed)
		}

		leaderboard = append(leaderboard, models.LeaderboardEntry{
			Rank:          rank,
			Username:      user.Username,
			TotalCoins:    user.TotalCoins,
			GamesPlayed:   user.GamesPlayed,
			GamesWon:      user.GamesWon,
			WinRate:       winRate,
			CurrentStreak: user.CurrentStreak,
		})
		rank++
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating leaderboard rows: %v", err)
	}

	return leaderboard, nil
}
