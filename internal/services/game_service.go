package services

import (
	"database/sql"
	"fmt"
	"rockpaperscissors/internal/models"
)

// GameService struct -> handles game logic and user interactions

type GameService struct {
	db          *sql.DB
	gameLogic   *GameLogicService
	userService *UserService
}

// creates a new game service
func NewGameService(db *sql.DB) *GameService {
	return &GameService{
		db:          db,
		gameLogic:   NewGameLogicService(),
		userService: NewUserService(db),
	}
}

// PlayGame -> handles the game logic and user interactions
// (g *GameService) -> pointer to the GameService struct
// (username string, playerChoice models.Choice) -> username and player choice is passed as parameters
// (*models.PlayGameResponse, error) -> return type and error
func (g *GameService) PlayGame(username string, playerChoice models.Choice) (*models.PlayGameResponse, error) {
	user, err := g.userService.GetUser(username)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	// game logic
	computerChoice := g.gameLogic.GenerateComputerChoice()
	result := g.gameLogic.DetermineWinner(playerChoice, computerChoice)
	coinsEarned := g.gameLogic.CalculateCoinsEarned(result, user.CurrentStreak)
	streakMultiplier := g.gameLogic.CalculateStreakMultiplier(user.CurrentStreak)
	newStreak := g.gameLogic.CalculateNewStreak(user.CurrentStreak, result)

	// update user stats
	newTotalCoins := user.TotalCoins + coinsEarned
	newGamesPlayed := user.GamesPlayed + 1
	newGamesWon := user.GamesWon
	if result == models.Win {
		newGamesWon++
	}
	// error handling
	err = g.userService.UpdateUserStats(user.ID, newTotalCoins, newStreak, newGamesPlayed, newGamesWon)
	if err != nil {
		return nil, fmt.Errorf("failed to update user stats: %v for user: %s", err, username)
	}

	// Save game record to database
	err = g.SaveGameRecord(user.ID, playerChoice, computerChoice, result, coinsEarned, streakMultiplier)
	if err != nil {
		return nil, fmt.Errorf("failed to save game record: %v", err)
	}

	// create response
	message := g.gameLogic.GetResultMessage(playerChoice, computerChoice, result, coinsEarned)

	return &models.PlayGameResponse{
		PlayerChoice:     playerChoice,
		ComputerChoice:   computerChoice,
		Result:           result,
		CoinsEarned:      coinsEarned,
		StreakMultiplier: streakMultiplier,
		NewStreak:        newStreak,
		TotalCoins:       newTotalCoins,
		Message:          message,
	}, nil
}

// SaveGameRecord saves an individual game record to the database
func (g *GameService) SaveGameRecord(userID int, playerChoice, computerChoice models.Choice, result models.GameResult, coinsEarned, streakMultiplier int) error {
	query := `
		INSERT INTO games (user_id, player_choice, computer_choice, result, coins_earned, streak_multiplier, played_at)
		VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
	`

	_, err := g.db.Exec(query, userID, string(playerChoice), string(computerChoice), string(result), coinsEarned, streakMultiplier)
	if err != nil {
		return fmt.Errorf("failed to insert game record: %v", err)
	}

	return nil
}

// GetUserGameHistory retrieves the game history for a specific user
func (g *GameService) GetUserGameHistory(username string, limit int) ([]models.Game, error) {
	if limit <= 0 {
		limit = 20 // Default to last 20 games
	}

	// First get the user to get their ID
	user, err := g.userService.GetUser(username)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	query := `
		SELECT id, user_id, player_choice, computer_choice, result, coins_earned, streak_multiplier, played_at
		FROM games 
		WHERE user_id = ?
		ORDER BY played_at DESC
		LIMIT ?
	`

	rows, err := g.db.Query(query, user.ID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query game history: %v", err)
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var game models.Game
		var playerChoice, computerChoice, result string

		err := rows.Scan(
			&game.ID,
			&game.UserID,
			&playerChoice,
			&computerChoice,
			&result,
			&game.CoinsEarned,
			&game.StreakMultiplier,
			&game.PlayedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan game row: %v", err)
		}

		// Convert string fields back to typed fields
		game.PlayerChoice = models.Choice(playerChoice)
		game.ComputerChoice = models.Choice(computerChoice)
		game.Result = models.GameResult(result)

		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating game rows: %v", err)
	}

	return games, nil
}
