package player

import (
	"GoDungeon/floor"
	"errors"
	"fmt"
)

type Player struct {
	x, y   int
	symbol rune
}

func NewPlayer(floor *floor.DungeonFloor) (*Player, error) {
	x, y := 1, 1
	if !floor.IsWalkable(x, y) {
		return nil, fmt.Errorf("starting position (%d%d) is not walkable", x, y)
	}
	return &Player{
		x:      x,
		y:      y,
		symbol: '@',
	}, nil
}

func (p *Player) MovePlayer(direction string, floor *floor.DungeonFloor) error {
	newX, newY := p.x, p.y

	switch direction {
	case "up":
		newY--
	case "down":
		newY++
	case "left":
		newX--
	case "right":
		newX++
	default:
		return fmt.Errorf("invalid direction: %s", direction)
	}

	if !floor.IsWalkable(newX, newY) {
		return errors.New("cannot move into wall or out of bounds")
	}

	p.x, p.y = newX, newY
	return nil
}

func (p *Player) GetPosition() (int, int, rune) {
	return p.x, p.y, p.symbol
}
