package services

import (
	"math/rand"
	"rockpaperscissors/internal/models"
	"time"
)

// GameLogicService handles the core game rules and logic
type GameLogicService struct {
	rng *rand.Rand
}

// NewGameLogicService creates a new game logic service
func NewGameLogicService() *GameLogicService {
	// Create a new random source with current time as seed
	source := rand.NewSource(time.Now().UnixNano())
	return &GameLogicService{
		rng: rand.New(source),
	}
}

// GenerateComputerChoice randomly selects rock, paper, or scissors
func (g *GameLogicService) GenerateComputerChoice() models.Choice {
	choices := []models.Choice{models.Rock, models.Paper, models.Scissors}
	randomIndex := g.rng.Intn(len(choices))
	return choices[randomIndex]
}

// determine the winner depending on who player and computer choice
func (g *GameLogicService) DetermineWinner(playerChoice, computerChoice models.Choice) models.GameResult {
	// Handle tie cases
	if playerChoice == computerChoice {
		return models.Tie
	}
	// switch logic to determine winner
	switch playerChoice {
	case models.Rock: // player plays rock, if computer plays scissors, player wins
		if computerChoice == models.Scissors {
			return models.Win
		}
	case models.Paper:
		if computerChoice == models.Rock {
			return models.Win
		}
	case models.Scissors:
		if computerChoice == models.Paper {
			return models.Win
		}
	}
	return models.Lose
}

// streak multiplier logic with a cap of 5
func (g *GameLogicService) CalculateStreakMultiplier(currentStreak int) int {
	switch {
	case currentStreak == 0:
		return 1 //return the reward multiplier
	case currentStreak == 1:
		return 2
	case currentStreak == 2:
		return 3
	case currentStreak == 3:
		return 4
	default:
		return 5 // max multiplier is 5
	}
}

// calculate the new streak after a game ends

func (g *GameLogicService) CalculateNewStreak(currentStreak int, result models.GameResult) int {
	if result == models.Win {
		return currentStreak + 1
	} else if result == models.Lose {
		return 0
	} else if result == models.Tie {
		return currentStreak
	}
	return 0 // default
}
