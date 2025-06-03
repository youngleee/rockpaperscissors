package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database connection
func InitDB() (*sql.DB, error) {
	// Ensure data directory exists
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %v", err)
	}

	// Database file path
	dbPath := filepath.Join(dataDir, "rockpaperscissors.db")

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Enable foreign key constraints
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to enable foreign keys: %v", err)
	}

	return db, nil
}

// RunMigrations executes database migrations
func RunMigrations(db *sql.DB) error {
	// Create users table
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		total_coins INTEGER DEFAULT 0,
		current_streak INTEGER DEFAULT 0,
		games_played INTEGER DEFAULT 0,
		games_won INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	// Create games table for game history
	gamesTable := `
	CREATE TABLE IF NOT EXISTS games (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		player_choice TEXT NOT NULL,
		computer_choice TEXT NOT NULL,
		result TEXT NOT NULL, -- 'win', 'lose', 'tie'
		coins_earned INTEGER DEFAULT 0,
		streak_multiplier INTEGER DEFAULT 1,
		played_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	// Create indexes for better performance
	indexesSQL := []string{
		"CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);",
		"CREATE INDEX IF NOT EXISTS idx_games_user_id ON games(user_id);",
		"CREATE INDEX IF NOT EXISTS idx_games_played_at ON games(played_at);",
		"CREATE INDEX IF NOT EXISTS idx_users_total_coins ON users(total_coins);",
	}

	// Execute migrations
	migrations := []string{usersTable, gamesTable}
	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("failed to execute migration: %v", err)
		}
	}

	// Create indexes
	for _, indexSQL := range indexesSQL {
		if _, err := db.Exec(indexSQL); err != nil {
			return fmt.Errorf("failed to create index: %v", err)
		}
	}

	return nil
} 