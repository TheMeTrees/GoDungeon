package floor

import (
	"testing"
)

func TestNewDungeonFloor(t *testing.T) {
	floor, err := NewDungeonFloor()
	if err != nil {
		t.Fatalf("NewDungeonFloor() error: %v", err)
	}
	if floor.Width != 20 || floor.Height != 10 {
		t.Errorf("Expected 20x10 floor, got %dx%d", floor.Width, floor.Height)
	}
	if floor.grid[0][0] != '#' || floor.grid[0][19] != '#' || floor.grid[9][0] != '#' || floor.grid[9][19] != '#' {
		t.Errorf("Edge tiles are not walls")
	}
	if floor.grid[1][1] != '_' {
		t.Errorf("Starting position (1,1) is not walkable")
	}
	hasWalkable := false
	for i := 0; i < floor.Height; i++ {
		for j := 0; j < floor.Width; j++ {
			if floor.grid[i][j] == '_' {
				hasWalkable = true
				break
			}
		}
	}
	if !hasWalkable {
		t.Error("Map has no walkable tiles")
	}
}

func TestIsWalkable(t *testing.T) {
	floor, err := NewDungeonFloor()
	if err != nil {
		t.Fatalf("NewDungeonFloor() error: %v", err)
	}

	tests := []struct {
		x, y     int
		expected bool
	}{
		{1, 1, true},    // Ensured walkable by NewDungeonFloor
		{0, 0, false},   // Wall (edge)
		{-1, 0, false},  // Out of bounds
		{10, 10, false}, // Out of bounds
	}

	for _, test := range tests {
		result := floor.IsWalkable(test.x, test.y)
		if result != test.expected {
			t.Errorf("IsWalkable(%d, %d) = %v; expected %v", test.x, test.y, result, test.expected)
		}
	}
}
