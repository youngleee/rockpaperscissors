package services

import (
	"fmt"
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

// calculate coins earned based on game result and current streak
func (g *GameLogicService) CalculateCoinsEarned(result models.GameResult, currentStreak int) int {
	// Only wins earn coins
	if result != models.Win {
		return 0
	}

	// Get the multiplier for the CURRENT streak (not new streak)
	multiplier := g.CalculateStreakMultiplier(currentStreak)

	// Calculate coins: base coins (10) * multiplier
	const baseCoins = 10
	return baseCoins * multiplier
}

func (g *GameLogicService) GetBeatMessage(winner, loser models.Choice) string {
	switch {
	case winner == models.Rock && loser == models.Scissors:
		return "Rock crushes Scissors!"
	case winner == models.Paper && loser == models.Rock:
		return "Paper smothers Rock!"
	case winner == models.Scissors && loser == models.Paper:
		return "Scissors snips Paper!"
	default:
		return ""
	}
}

func (g *GameLogicService) GetResultMessage(playerChoice, computerChoice models.Choice, result models.GameResult, coinsEarned int) string {
	var baseMessage string = fmt.Sprintf("You chose %s, computer chose %s. ", playerChoice, computerChoice) // cool syntax for string interpolation!

	switch result {
	case models.Win:
		var beatMessage string = g.GetBeatMessage(playerChoice, computerChoice)
		return fmt.Sprintf("%s%s You won! +%d coins", baseMessage, beatMessage, coinsEarned)
	case models.Lose:
		var beatMessage string = g.GetBeatMessage(computerChoice, playerChoice) // Computer beat player
		return fmt.Sprintf("%s%s You lost!", baseMessage, beatMessage)
	case models.Tie:
		return fmt.Sprintf("%s It's a tie! No coins earned, but streak preserved.", baseMessage)
	default:
		return baseMessage + "Unknown Result/Error"
	}
}
