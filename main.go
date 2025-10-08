package main

import (
	"GoDungeon/floor"
	"GoDungeon/player"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

func main() {
	floor, err := floor.NewDungeonFloor()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing floor: %v\n", err)
		os.Exit(1)
	}

	player, err := player.NewPlayer(floor)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing player: %v\n", err)
		os.Exit(1)
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting terminal to raw mode: %v\n", err)
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	buf := make([]byte, 1)

	for {
		// Render the map with the player
		x, y, symbol := player.GetPosition()
		if err := floor.Render(x, y, symbol); err != nil {
			fmt.Fprintf(os.Stderr, "Error rendering map: %v\n", err)
			os.Exit(1)
		}

		// Prompt for input
		fmt.Print("Move (WASD) or Q to quit: ")

		// Read a single keypress
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		input := strings.ToLower(string(buf[0])) // Convert to lowercase

		// Process input
		switch input {
		case "w":
			if err := player.MovePlayer("up", floor); err != nil {
				fmt.Printf("Cannot move up: %v\n", err)
			}
		case "a":
			if err := player.MovePlayer("left", floor); err != nil {
				fmt.Printf("Cannot move left: %v\n", err)
			}
		case "s":
			if err := player.MovePlayer("down", floor); err != nil {
				fmt.Printf("Cannot move down: %v\n", err)
			}
		case "d":
			if err := player.MovePlayer("right", floor); err != nil {
				fmt.Printf("Cannot move right: %v\n", err)
			}
		case "q":
			fmt.Println("Exiting game...")
			os.Exit(0)
		default:
			fmt.Printf("Invalid input: %q. Use WASD to move or Q to quit.\n", input)
		}
	}
}
