package floor

import (
	"fmt"
	"os"
	"os/exec"
)

type DungeonFloor struct {
	grid   [][]rune
	Width  int
	Height int
}

func NewDungeonFloor() (*DungeonFloor, error) {
	width, height := 20, 10
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

	// First room vertical wall
	grid[3][1] = '#'
	grid[3][2] = '#'
	grid[3][3] = '#'
	grid[3][4] = '#'

	// First room horizontal wall
	grid[1][6] = '#'
	grid[2][6] = '#'
	grid[3][6] = '#'

	// Second room vertical wall
	grid[4][9] = '#'
	grid[4][10] = '#'
	grid[4][11] = '#'
	grid[4][12] = '#'
	grid[4][13] = '#'
	grid[4][14] = '#'
	grid[4][15] = '#'
	grid[4][16] = '#'
	grid[4][17] = '#'
	grid[4][18] = '#'
	grid[4][19] = '#'

	// Second room horizontal wall
	grid[5][9] = '#'
	grid[6][9] = '#'
	grid[8][9] = '#'

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
