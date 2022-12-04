package challenges

import (
	"strconv"
	"strings"
)

type numRange struct {
	min int
	max int
}

func CountOfFullyContainedRanges(input string) int {
	lines := strings.Split(input, "\n")

	var total int
	for _, ln := range lines {
		parts := strings.Split(ln, ",")
		r1 := strings.Split(parts[0], "-")
		r2 := strings.Split(parts[1], "-")

		first := parseMinMaxRange(r1)
		second := parseMinMaxRange(r2)

		firstDiff := first.max - first.min
		secondDiff := second.max - second.min

		pass := true
		// First has a larger range
		if firstDiff > secondDiff {
			if first.min > second.min {
				pass = false
			}

			if first.max < second.max {
				pass = false
			}
		} else if secondDiff >= firstDiff {
			if second.min > first.min {
				pass = false
			}

			if second.max < first.max {
				pass = false
			}
		}

		if !pass {
			continue
		}

		total++
	}

	return total
}

func CountPairsOfOverlap(input string) int {
	lines := strings.Split(input, "\n")

	var total int
	for _, ln := range lines {
		parts := strings.Split(ln, ",")
		r1 := strings.Split(parts[0], "-")
		r2 := strings.Split(parts[1], "-")

		first := parseMinMaxRange(r1)
		second := parseMinMaxRange(r2)

		overlap := first.max - second.min
		if second.min <= first.min {
			overlap = second.max - first.min
		}

		if overlap >= 0 {
			total++
		}
	}

	return total
}

func parseMinMaxRange(s []string) numRange {
	min, _ := strconv.ParseInt(s[0], 10, 64)
	max, _ := strconv.ParseInt(s[1], 10, 64)
	return numRange{
		min: int(min),
		max: int(max),
	}
}
