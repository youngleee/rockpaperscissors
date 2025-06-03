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
