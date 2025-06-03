package services

import (
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
}
