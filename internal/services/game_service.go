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
