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

	// Test to check if streak multiplier is applied correctly
	t.Run("CalculateStreakMultiplier", func(t *testing.T) {
		multiplier := gameLogic.CalculateStreakMultiplier(0)
		if multiplier != 1 {
			t.Errorf("Streak: 0, Expected multiplier: 1, got %d.", multiplier)
		}
		multiplier = gameLogic.CalculateStreakMultiplier(1)
		if multiplier != 2 {
			t.Errorf("Streak: 1, Expected multiplier: 2, got %d.", multiplier)
		}
		multiplier = gameLogic.CalculateStreakMultiplier(2)
		if multiplier != 3 {
			t.Errorf("Streak: 2, Expected multiplier: 3, got %d.", multiplier)
		}
		multiplier = gameLogic.CalculateStreakMultiplier(3)
		if multiplier != 4 {
			t.Errorf("Streak: 3, Expected multiplier: 4, got %d.", multiplier)
		}
		multiplier = gameLogic.CalculateStreakMultiplier(4)
		if multiplier != 5 {
			t.Errorf("Streak: 4, Expected multiplier: 5, got %d.", multiplier)
		}
	})

	t.Run("GameScenario", func(t *testing.T) {
		currentStreak := 1
		result := gameLogic.DetermineWinner(models.Rock, models.Scissors)
		coinsEarned := gameLogic.CalculateCoinsEarned(result, currentStreak)
		newStreak := gameLogic.CalculateNewStreak(currentStreak, result)
		resultMessage := gameLogic.GetResultMessage(models.Rock, models.Scissors, result, coinsEarned)

		// Verify game result
		if result != models.Win {
			t.Errorf("Expected win, got %s.", result)
		}

		// Verify coins earned (current streak 1 = 2x multiplier = 10 * 2 = 20 coins)
		if coinsEarned != 20 {
			t.Errorf("Expected 20 coins, got %d.", coinsEarned)
		}

		// Verify new streak (1 + 1 = 2)
		if newStreak != 2 {
			t.Errorf("Expected new streak: 2, got %d.", newStreak)
		}

		// Verify result message matches what our function actually returns
		expectedMessage := "You chose rock, computer chose scissors. Rock crushes Scissors! You won! +20 coins"
		if resultMessage != expectedMessage {
			t.Errorf("Expected: '%s', got: '%s'", expectedMessage, resultMessage)
		}

		// Log the actual values for debugging
		t.Logf("Current streak: %d", currentStreak)
		t.Logf("Result: %s", result)
		t.Logf("Coins earned: %d", coinsEarned)
		t.Logf("New streak: %d", newStreak)
		t.Logf("Message: %s", resultMessage)
	})
}
