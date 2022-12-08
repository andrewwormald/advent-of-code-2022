package challenges

import (
	"fmt"
	"strconv"
	"strings"
)

type data struct {
	trees     [][]int
	positions map[string]int

	// visible maps position to height
	visible map[string]bool
	scores  map[string]int
}

func TreeTop(input string) (int, int) {
	parts := strings.Split(input, "")

	var (
		d   data
		cur []int
	)

	for _, ln := range parts {
		if ln == "\n" {
			if len(cur) == 0 {
				continue
			}

			d.trees = append(d.trees, cur)
			cur = []int{}
			continue
		}

		height, _ := strconv.Atoi(ln)
		cur = append(cur, height)
	}

	return part1(d), part2(d)
}

func part1(d data) int {
	d.visible = make(map[string]bool)
	d.positions = make(map[string]int)
	for y, row := range d.trees {
		for x, tree := range row {
			d.positions[pos(x, y)] = tree
		}
	}

	for y, row := range d.trees {
		for x, tree := range row {
			position := pos(x, y)

			// Add top and bottom rows as visible
			if x == 0 || x == len(d.trees[0])-1 {
				d.visible[pos(x, y)] = true
				continue
			}

			// Add side edges as visible
			if y == 0 || y == len(d.trees)-1 {
				d.visible[pos(x, y)] = true
				continue
			}

			// Look left to right for right side consideration
			var lrMax int
			for i := x + 1; i <= len(row)-1; i++ {
				height := d.positions[pos(i, y)]
				if height > lrMax {
					lrMax = height
				}
			}

			if lrMax < tree {
				d.visible[position] = true
			}

			// Look left to right for right side consideration
			var rlMax int
			for i := x - 1; i >= 0; i-- {
				height := d.positions[pos(i, y)]
				if height > rlMax {
					rlMax = height
				}
			}

			if rlMax < tree {
				d.visible[position] = true
			}

			// Look top down
			var tdMax int
			for i := y + 1; i <= len(d.trees)-1; i++ {
				height := d.positions[pos(x, i)]
				if height > tdMax {
					tdMax = height
				}
			}

			if tdMax < tree {
				d.visible[position] = true
			}

			// Look left to right for right side consideration
			var dtMax int
			for i := y - 1; i >= 0; i-- {
				height := d.positions[pos(x, i)]
				if height > dtMax {
					dtMax = height
				}
			}

			if dtMax < tree {
				d.visible[position] = true
			}
		}
	}

	var totalVisible int
	for _, isVisible := range d.visible {
		if isVisible {
			totalVisible++
		}
	}

	return totalVisible
}

func part2(d data) int {
	d.scores = make(map[string]int)
	d.positions = make(map[string]int)

	for y, row := range d.trees {
		for x, tree := range row {
			d.positions[pos(x, y)] = tree
		}
	}

	for y, row := range d.trees {
		for x, tree := range row {
			position := pos(x, y)

			// Exclude edges
			if x == 0 || x == len(d.trees[0])-1 {
				continue
			}

			// Exclude edges
			if y == 0 || y == len(d.trees)-1 {
				continue
			}

			// Look left to right for right side consideration
			var lrMaxDistance int
			for i := x + 1; i <= len(row)-1; i++ {
				height := d.positions[pos(i, y)]
				if height < tree {
					lrMaxDistance++
					continue
				}

				lrMaxDistance++
				break
			}

			// Look left to right for right side consideration
			var rlMaxDistance int
			for i := x - 1; i >= 0; i-- {
				height := d.positions[pos(i, y)]
				if height < tree {
					rlMaxDistance++
					continue
				}

				rlMaxDistance++
				break
			}

			// Look top down
			var tdMaxDistance int
			for i := y + 1; i <= len(d.trees)-1; i++ {
				height := d.positions[pos(x, i)]
				if height < tree {
					tdMaxDistance++
					continue
				}

				tdMaxDistance++
				break
			}

			// Look left to right for right side consideration
			var dtMaxDistance int
			for i := y - 1; i >= 0; i-- {
				height := d.positions[pos(x, i)]
				if height < tree {
					dtMaxDistance++
					continue
				}

				dtMaxDistance++
				break
			}

			d.scores[position] = lrMaxDistance * rlMaxDistance * tdMaxDistance * dtMaxDistance
		}
	}

	var max int
	for _, score := range d.scores {
		if score > max {
			max = score
		}
	}

	return max
}

func pos(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}
