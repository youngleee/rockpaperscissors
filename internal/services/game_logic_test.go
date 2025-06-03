package services

import (
	"rockpaperscissors/internal/models"
	"testing"
)

// TestGameLogicService tests our game logic functions
func TestGameLogicService(t *testing.T) {
	// Create a new game logic service to test
	gameLogic := NewGameLogicService()

	t.Run("GenerateComputerChoice", func(t *testing.T) {
		choice := gameLogic.GenerateComputerChoice()
		if !choice.IsValid() {
			t.Errorf("Invalid choice generated: %v", choice)
		}
		t.Logf("Computer chose: %s", choice)
	})

	// Test to determine the winner of a game
	t.Run("DetermineWinner", func(t *testing.T) {
		//test if rock beats scissors
		result := gameLogic.DetermineWinner(models.Rock, models.Scissors)
		if result != models.Win {
			t.Errorf("Rock should beat Scissors, Expected win, got %s.", result)
		}
		result = gameLogic.DetermineWinner(models.Scissors, models.Rock)
		if result != models.Lose {
			t.Errorf("Scissors should lose to Rock, Expected lose, got %s.", result)
		}
		result = gameLogic.DetermineWinner(models.Rock, models.Rock)
		if result != models.Tie {
			t.Errorf("Rock should tie with Rock, Expected tie, got %s.", result)
		}
	})
}
