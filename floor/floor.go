package floor

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type DungeonFloor struct {
	grid   [][]rune
	Width  int
	Height int
}

func NewDungeonFloor() (*DungeonFloor, error) {
	width, height := 20, 10
	rand.Seed(time.Now().UnixNano())

	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '_'
		}
	}

	for i := 0; i < height; i++ {
		grid[i][0] = '#'
		grid[i][width-1] = '#'
		for j := 1; j < width; j++ {
			grid[0][j] = '#'
			grid[height-1][j] = '#'
		}
	}

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if rand.Float64() < 0.3 {
				grid[i][j] = '#'
			}
		}
	}

	// Walkable starting position
	grid[1][1] = '_'

	return &DungeonFloor{grid, width, height}, nil
}

func (df *DungeonFloor) Render(playerX, playerY int, playerSymbol rune) error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error clearing console: %v", err)
	}

	for i := 0; i < df.Height; i++ {
		for j := 0; j < df.Width; j++ {
			if i == playerY && j == playerX {
				fmt.Print(string(playerSymbol))
			} else {
				fmt.Print(string(df.grid[i][j]))
			}
		}
		fmt.Println()
	}

	return nil
}

func (df *DungeonFloor) IsWalkable(x, y int) bool {
	if x < 0 || x >= df.Width || y < 0 || y >= df.Height {
		return false
	}

	return df.grid[y][x] == '_'
}
