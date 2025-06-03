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

func (g *GameLogicService) DetermineWinner(playerChoice, computerChoice models.Choice) models.GameResult {
	// Handle tie cases
	if playerChoice == computerChoice {
		return models.Tie
	}
	// switch logic to determine winner
	switch playerChoice {
	case models.Rock:
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
