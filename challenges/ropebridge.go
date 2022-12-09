package challenges

import (
	"fmt"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func (p *position) DistanceFrom(p2 position) (xDist int, yDist int) {
	return p.x - p2.x, p.y - p2.y
}

func (p *position) ClosestPositionTo(p2 position) position {
	xDist, yDist := p2.DistanceFrom(*p)

	if xDist >= -1 && xDist <= 1 && yDist >= -1 && yDist <= 1 {
		return *p
	}

	xMod := -1
	if xDist < 0 {
		xMod = 1
	}

	yMod := -1
	if yDist < 0 {
		yMod = 1
	}

	var newPos position
	if xDist < -1 || xDist > 1 {
		newPos.x = p.x + (xDist + xMod)
	} else {
		newPos.x = p2.x
	}

	if yDist < -1 || yDist > 1 {
		newPos.y = p.y + (yDist + yMod)
	} else {
		newPos.y = p2.y
	}

	return newPos
}

func (p *position) String() string {
	return fmt.Sprintf("%v-%v", p.x, p.y)
}

func (p *position) Move(d direction) {
	switch d {
	case Up:
		p.y++
	case Down:
		p.y--
	case Left:
		p.x--
	case Right:
		p.x++
	default:
		return
	}
}

func (p *position) CountMovesTo(p2 position) int {
	closest := p.ClosestPositionTo(p2)

	xDist, yDist := closest.DistanceFrom(*p)
	if xDist < 0 {
		xDist = xDist * -1
	}

	if yDist < 0 {
		yDist = yDist * -1
	}

	max := xDist
	if yDist > max {
		max = yDist
	}

	return max
}

type direction string

func parseDirection(s string) direction {
	switch s {
	case string(Up):
		return Up
	case string(Down):
		return Down
	case string(Left):
		return Left
	case string(Right):
		return Right
	default:
		return UnknownDirection
	}
}

const (
	UnknownDirection direction = ""
	Up               direction = "U"
	Down             direction = "D"
	Left             direction = "L"
	Right            direction = "R"
)

func RopeBridge(input string, count int) int {
	lines := strings.Split(input, "\n")

	head := &position{
		x: 1,
		y: 1,
	}

	positions := make(map[int]*position)
	tracker := make(map[int]map[string]position)
	for i := 1; i < count; i++ {
		pos := position{
			x: 1,
			y: 1,
		}
		positions[i] = &pos
		tracker[i] = map[string]position{}
	}

	for _, ln := range lines {
		if ln == "" {
			continue
		}

		dir, move, _ := strings.Cut(ln, " ")
		d := parseDirection(dir)
		m, _ := strconv.Atoi(move)

		for i := 0; i < m; i++ {
			head.Move(d)

			previous := *head
			for j := 1; j < count; j++ {
				*positions[j] = positions[j].ClosestPositionTo(previous)
				tracker[j][positions[j].String()] = *positions[j]
				previous = *positions[j]
			}
		}
	}

	return len(tracker[count-1])
}
