package challenges

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
)

func HillClimbing(input string) int {
	lines := strings.Split(input, "\n")

	var (
		grid    [][]string
		startP  position
		endP    position
		skipped int
	)
	for y, ln := range lines {
		if len(ln) == 0 {
			skipped++
			continue
		}

		chars := strings.Split(ln, "")
		var cList []string
		for x, char := range chars {
			if char == "S" {
				startP = position{
					x: x,
					y: y - skipped,
				}

				char = "a"
			}

			if char == "E" {
				endP = position{
					x: x,
					y: y - skipped,
				}

				char = "z"
			}

			cList = append(cList, char)
		}

		grid = append(grid, cList)
	}

	return aStarSearch(startP, endP, grid)
}

func aStarSearch(start, end position, grid [][]string) int {
	var (
		open  [][]int
		moves []position
	)
	closed := make(map[position]bool)

	for _, r := range grid {
		var row []int
		for _ = range r {
			row = append(row, 0)
		}

		open = append(open, row)
	}

	options := []position{
		{0, 1},  // One up
		{0, -1}, // One down
		{-1, 0}, // One to the left
		{1, 0},  // One to the right
	}

	cur := start
	open[start.y][start.x] = 1
	moves = append(moves, cur)

	for !cur.EqualTo(end) {
		// Evaluate up, down, left, right
		var potential []position
		for _, opt := range options {
			x := cur.x + opt.x
			y := cur.y + opt.y

			if x < 0 {
				continue
			} else if y < 0 {
				continue
			}

			if y > len(grid)-1 {
				continue
			}

			if x > len(grid[y])-1 {
				continue
			}

			letter := grid[y][x]
			curLetter := grid[cur.y][cur.x]

			diff := CharPosition(letter) - CharPosition(curLetter)
			if diff <= 1 {
				if open[y][x] == 1 {
					continue
				}

				pp := position{
					x: x,
					y: y,
				}

				if closed[pp] {
					continue
				}

				potential = append(potential, pp)
			}
		}

		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()

		print(start, end, cur, open)

		time.Sleep(time.Second / 1000)

		if len(potential) == 0 {
			// Move backwards and record bad nodes
			closed[cur] = true
			open[cur.y][cur.x] = 0
			moves = remove(moves, len(moves)-1)
			cur = moves[len(moves)-1]
			continue
		}

		min := math.MaxInt
		var minP position
		for _, p := range potential {
			h := p.DistanceTo(end)
			g := len(moves) - 1
			f := g + h

			if f <= min {
				min = f
				minP = p
			}
		}

		cur = minP
		open[cur.y][cur.x] = 1
		moves = append(moves, cur)
	}
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	print(start, end, cur, open)

	return len(moves) - 1
}

func print(start, end, cur position, open [][]int) {
	fmt.Println(start, end, cur.DistanceTo(end))
	for y, row := range open {
		for x, v := range row {
			if start.y == y && start.x == x {
				fmt.Print("S")
				continue
			}

			if end.y == y && end.x == x {
				fmt.Print("E")
				continue
			}

			if v == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func remove(slice []position, s int) []position {
	return append(slice[:s], slice[s+1:]...)
}
