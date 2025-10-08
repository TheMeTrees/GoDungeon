package player

import (
	"GoDungeon/floor"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	floor, err := floor.NewDungeonFloor()
	if err != nil {
		t.Errorf("NewDungeonFloor() error: %v", err)
	}

	player, err := NewPlayer(floor)
	if err != nil {
		t.Errorf("NewPlayer() error: %v", err)
	}
	if player.x != 1 || player.y != 1 {
		t.Errorf("Expected player at (1,1), got (%d, %d)", player.x, player.y)
	}
	if player.symbol != '@' {
		t.Errorf("Expected player symbol ('@'), got (%c)", player.symbol)
	}
}

func TestMovePlayer(t *testing.T) {
	floor, err := floor.NewDungeonFloor()
	if err != nil {
		t.Fatalf("NewDungeonFloor() error: %v", err)
	}
	player, err := NewPlayer(floor)
	if err != nil {
		t.Fatalf("NewPlayer() error: %v", err)
	}

	tests := []struct {
		direction  string
		newX, newY int
		shouldFail bool
	}{
		{"right", 2, 1, false},  // Valid move
		{"up", 1, 1, true},      // Blocked by top wall
		{"down", 1, 2, false},   // Valid move
		{"left", 1, 1, true},    // Blocked by left wall
		{"invalid", 1, 1, true}, // Invalid direction
	}

	for _, test := range tests {
		originalX, originalY := player.x, player.y
		err := player.MovePlayer(test.direction, floor)
		if test.shouldFail {
			if err == nil {
				t.Errorf("MovePlayer(%s) should have failed", test.direction)
			}
		} else {
			if err != nil {
				t.Errorf("MovePlayer(%s) failed: %v", test.direction, err)
			}
			if player.x != test.newX || player.y != test.newY {
				t.Errorf("MovePlayer(%s) = (%d,%d); want (%d,%d)", test.direction, player.x, player.y, test.newX, test.newY)
			}
		}

		player.x, player.y = originalX, originalY
	}
}
