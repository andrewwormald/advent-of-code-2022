package challenges

import (
	"strconv"
	"strings"
)

func SupplyStacksPart1(stack string, moves string) string {
	st := parseStack(stack)

	for _, move := range parseMoves(moves) {
		st.Move(move, true)
	}

	var tops string
	for i := 1; i <= 9; i++ {
		temp := strings.ReplaceAll(st[i][0], "[", "")
		tops += strings.ReplaceAll(temp, "]", "")
	}

	return tops
}

type move struct {
	Count int
	From  int
	To    int
}

func parseMoves(s string) []move {
	rows := strings.Split(s, "\n")
	var moves []move
	for _, row := range rows {
		if row == "" {
			continue
		}

		moves = append(moves, parseMove(row))
	}

	return moves
}

func parseMove(s string) move {
	s = strings.Replace(s, "move ", "", 1)
	s = strings.Replace(s, " from ", "-", 1)
	s = strings.Replace(s, " to ", "-", 1)
	parts := strings.Split(s, "-")

	c, _ := strconv.Atoi(parts[0])
	f, _ := strconv.Atoi(parts[1])
	t, _ := strconv.Atoi(parts[2])

	return move{
		Count: c,
		From:  f,
		To:    t,
	}
}

type stack map[int][]string

func (s stack) Move(m move, multipleAtOnce bool) {
	col := s[m.From]
	var taken []string
	for i, box := range col {
		if len(taken) == m.Count {
			break
		}

		if box == emptySpace {
			continue
		}

		taken = append(taken, box)
		s[m.From][i] = emptySpace
	}

	var reordered []string
	if !multipleAtOnce {
		for i := len(taken) - 1; i >= 0; i-- {
			reordered = append(reordered, taken[i])
		}
	} else {
		reordered = taken
	}

	var filtered []string
	for _, s := range s[m.From] {
		if s == emptySpace {
			continue
		}

		filtered = append(filtered, s)
	}

	s[m.From] = filtered

	newCol := reordered
	for _, box := range s[m.To] {
		if box == emptySpace {
			continue
		}

		newCol = append(newCol, box)
	}

	s[m.To] = newCol
}

var emptySpace = "[|]"

func parseStack(s string) stack {
	rows := strings.Split(s, "\n")
	indexRow := rows[len(rows)-1]

	colCount := len(strings.ReplaceAll(indexRow, " ", ""))

	st := make(map[int][]string)

	for i := 1; i <= colCount; i++ {
		st[i] = []string{}
	}

	var newRows [][]string
	for i, row := range rows {
		if i == len(rows)-1 {
			continue
		}

		var newRow []string
		var char string
		for i, ch := range row {
			if (i+1)%4 == 0 {
				if char == "   " {
					char = emptySpace
				}
				newRow = append(newRow, char)
				char = ""
				continue
			}

			char += string(ch)
		}

		// Ensure remaining characters are used
		if char != "" {
			if char == "   " {
				char = emptySpace
			}
			newRow = append(newRow, char)
		}

		newRows = append(newRows, newRow)
	}

	for _, row := range newRows {
		for colIndex, col := range row {
			st[colIndex+1] = append(st[colIndex+1], col)
		}
	}

	return st
}

var TestStack = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3`

var TestMoves = `
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

var SupplyStacks = `            [G] [W]         [Q]    
[Z]         [Q] [M]     [J] [F]    
[V]         [V] [S] [F] [N] [R]    
[T]         [F] [C] [H] [F] [W] [P]
[B] [L]     [L] [J] [C] [V] [D] [V]
[J] [V] [F] [N] [T] [T] [C] [Z] [W]
[G] [R] [Q] [H] [Q] [W] [Z] [G] [B]
[R] [J] [S] [Z] [R] [S] [D] [L] [J]
 1   2   3   4   5   6   7   8   9`

var SupplyStackMoves = `
move 6 from 5 to 7
move 2 from 9 to 1
move 4 from 8 to 6
move 1 from 8 to 1
move 2 from 9 to 1
move 1 from 6 to 1
move 13 from 7 to 8
move 1 from 2 to 8
move 9 from 1 to 5
move 1 from 3 to 8
move 3 from 6 to 7
move 4 from 4 to 1
move 11 from 5 to 6
move 6 from 6 to 9
move 3 from 4 to 2
move 7 from 8 to 6
move 1 from 7 to 5
move 1 from 4 to 3
move 7 from 1 to 5
move 2 from 2 to 7
move 4 from 9 to 6
move 1 from 3 to 6
move 1 from 1 to 9
move 1 from 3 to 6
move 1 from 5 to 8
move 4 from 6 to 7
move 3 from 8 to 7
move 7 from 5 to 7
move 1 from 3 to 1
move 1 from 2 to 6
move 14 from 6 to 5
move 2 from 5 to 2
move 3 from 9 to 2
move 6 from 2 to 9
move 7 from 8 to 6
move 7 from 7 to 3
move 2 from 8 to 7
move 6 from 3 to 7
move 17 from 7 to 1
move 1 from 3 to 1
move 1 from 2 to 5
move 4 from 5 to 6
move 17 from 6 to 9
move 7 from 9 to 4
move 1 from 2 to 7
move 2 from 5 to 4
move 3 from 7 to 8
move 7 from 5 to 2
move 6 from 2 to 8
move 8 from 9 to 6
move 1 from 2 to 3
move 8 from 4 to 9
move 7 from 6 to 9
move 18 from 1 to 7
move 1 from 1 to 8
move 2 from 6 to 9
move 1 from 3 to 9
move 1 from 4 to 6
move 1 from 8 to 3
move 1 from 3 to 1
move 10 from 7 to 2
move 9 from 8 to 4
move 1 from 6 to 4
move 2 from 7 to 8
move 5 from 4 to 9
move 17 from 9 to 5
move 2 from 7 to 6
move 5 from 9 to 7
move 5 from 4 to 2
move 8 from 2 to 4
move 8 from 4 to 3
move 2 from 6 to 5
move 2 from 8 to 5
move 3 from 9 to 3
move 4 from 7 to 3
move 6 from 9 to 6
move 4 from 6 to 9
move 5 from 9 to 3
move 8 from 5 to 2
move 1 from 1 to 9
move 1 from 6 to 3
move 1 from 9 to 4
move 5 from 7 to 4
move 19 from 3 to 1
move 4 from 2 to 8
move 13 from 5 to 1
move 1 from 6 to 3
move 3 from 3 to 6
move 2 from 8 to 9
move 4 from 2 to 9
move 2 from 2 to 6
move 1 from 1 to 6
move 5 from 1 to 9
move 10 from 9 to 3
move 15 from 1 to 6
move 21 from 6 to 2
move 20 from 2 to 1
move 2 from 8 to 9
move 28 from 1 to 2
move 6 from 4 to 6
move 2 from 1 to 5
move 3 from 3 to 4
move 2 from 5 to 4
move 1 from 4 to 3
move 3 from 4 to 5
move 2 from 5 to 4
move 1 from 1 to 8
move 25 from 2 to 9
move 1 from 4 to 6
move 1 from 3 to 8
move 4 from 3 to 6
move 1 from 4 to 9
move 2 from 6 to 3
move 1 from 5 to 9
move 5 from 2 to 8
move 7 from 9 to 6
move 2 from 9 to 4
move 3 from 2 to 1
move 3 from 3 to 4
move 1 from 3 to 5
move 16 from 6 to 3
move 7 from 8 to 3
move 5 from 4 to 3
move 1 from 1 to 3
move 1 from 2 to 6
move 1 from 5 to 6
move 21 from 3 to 5
move 2 from 1 to 2
move 1 from 6 to 7
move 10 from 9 to 8
move 1 from 6 to 5
move 5 from 8 to 7
move 12 from 5 to 3
move 20 from 3 to 6
move 4 from 7 to 9
move 1 from 7 to 3
move 1 from 2 to 5
move 1 from 3 to 8
move 2 from 8 to 4
move 4 from 8 to 7
move 3 from 6 to 1
move 1 from 1 to 5
move 2 from 9 to 2
move 2 from 1 to 5
move 2 from 5 to 6
move 3 from 7 to 1
move 2 from 1 to 4
move 4 from 6 to 8
move 3 from 4 to 7
move 3 from 2 to 5
move 2 from 7 to 9
move 9 from 9 to 8
move 1 from 4 to 1
move 7 from 5 to 7
move 1 from 7 to 8
move 1 from 3 to 1
move 4 from 7 to 5
move 2 from 1 to 9
move 1 from 1 to 2
move 5 from 5 to 4
move 1 from 2 to 6
move 5 from 7 to 9
move 5 from 4 to 7
move 11 from 9 to 6
move 14 from 8 to 9
move 23 from 6 to 5
move 6 from 9 to 5
move 1 from 6 to 2
move 10 from 5 to 3
move 1 from 4 to 9
move 1 from 2 to 1
move 2 from 7 to 3
move 10 from 5 to 7
move 8 from 5 to 2
move 5 from 3 to 5
move 7 from 5 to 8
move 1 from 2 to 7
move 9 from 7 to 9
move 3 from 2 to 3
move 2 from 6 to 2
move 2 from 3 to 6
move 4 from 7 to 5
move 1 from 1 to 5
move 4 from 3 to 1
move 2 from 5 to 2
move 1 from 3 to 2
move 2 from 6 to 8
move 7 from 5 to 3
move 9 from 2 to 4
move 2 from 1 to 2
move 2 from 5 to 3
move 1 from 4 to 9
move 1 from 6 to 9
move 1 from 4 to 2
move 2 from 1 to 7
move 3 from 2 to 6
move 4 from 8 to 7
move 2 from 8 to 3
move 2 from 3 to 7
move 1 from 6 to 5
move 2 from 8 to 2
move 5 from 4 to 1
move 8 from 9 to 8
move 1 from 5 to 7
move 10 from 9 to 2
move 8 from 8 to 2
move 1 from 1 to 6
move 12 from 3 to 9
move 7 from 7 to 4
move 13 from 2 to 4
move 7 from 2 to 7
move 1 from 6 to 7
move 3 from 9 to 8
move 2 from 6 to 3
move 1 from 3 to 2
move 1 from 3 to 9
move 3 from 1 to 5
move 1 from 1 to 6
move 4 from 7 to 6
move 5 from 7 to 1
move 1 from 2 to 1
move 6 from 9 to 4
move 5 from 9 to 7
move 3 from 8 to 3
move 22 from 4 to 9
move 24 from 9 to 8
move 1 from 9 to 2
move 2 from 4 to 3
move 10 from 8 to 3
move 1 from 2 to 1
move 1 from 3 to 8
move 1 from 6 to 3
move 1 from 1 to 4
move 4 from 3 to 4
move 4 from 6 to 1
move 2 from 4 to 5
move 4 from 7 to 2
move 7 from 4 to 6
move 4 from 6 to 1
move 2 from 6 to 3
move 1 from 6 to 2
move 5 from 5 to 2
move 12 from 3 to 5
move 3 from 7 to 8
move 6 from 2 to 3
move 11 from 1 to 9
move 1 from 1 to 7
move 1 from 7 to 5
move 2 from 3 to 9
move 2 from 9 to 7
move 4 from 2 to 5
move 2 from 7 to 1
move 17 from 8 to 1
move 1 from 3 to 2
move 16 from 1 to 3
move 8 from 3 to 4
move 2 from 8 to 3
move 2 from 1 to 5
move 1 from 2 to 6
move 12 from 5 to 8
move 1 from 6 to 3
move 9 from 3 to 9
move 8 from 4 to 6
move 2 from 1 to 6
move 6 from 8 to 4
move 3 from 4 to 6
move 1 from 1 to 9
move 11 from 6 to 8
move 3 from 4 to 3
move 17 from 9 to 5
move 2 from 6 to 7
move 1 from 9 to 1
move 2 from 8 to 6
move 1 from 7 to 5
move 1 from 8 to 9
move 1 from 1 to 7
move 3 from 9 to 6
move 2 from 7 to 8
move 1 from 9 to 6
move 15 from 5 to 2
move 9 from 3 to 9
move 11 from 8 to 3
move 6 from 9 to 8
move 4 from 6 to 7
move 3 from 3 to 7
move 5 from 5 to 6
move 7 from 7 to 5
move 3 from 6 to 1
move 2 from 1 to 4
move 1 from 9 to 2
move 2 from 9 to 3
move 2 from 6 to 3
move 1 from 1 to 8
move 6 from 5 to 9
move 8 from 2 to 5
move 10 from 8 to 5
move 1 from 2 to 9
move 21 from 5 to 9
move 2 from 8 to 4
move 5 from 9 to 1
move 2 from 5 to 2
move 15 from 9 to 2
move 1 from 5 to 9
move 9 from 9 to 3
move 1 from 1 to 6
move 3 from 4 to 1
move 20 from 3 to 5
move 20 from 5 to 4
move 7 from 4 to 3
move 1 from 1 to 7
move 11 from 4 to 5
move 4 from 3 to 2
move 11 from 5 to 4
move 2 from 6 to 7
move 4 from 3 to 9
move 2 from 2 to 8
move 2 from 9 to 4
move 6 from 4 to 6
move 2 from 7 to 9
move 1 from 7 to 6
move 1 from 4 to 9
move 4 from 4 to 6
move 2 from 8 to 6
move 1 from 4 to 3
move 1 from 4 to 6
move 1 from 3 to 1
move 3 from 4 to 3
move 9 from 2 to 8
move 2 from 3 to 7
move 5 from 6 to 2
move 2 from 7 to 5
move 1 from 5 to 2
move 1 from 9 to 3
move 1 from 5 to 1
move 13 from 2 to 5
move 4 from 9 to 5
move 1 from 3 to 4
move 9 from 2 to 3
move 7 from 3 to 2
move 11 from 5 to 6
move 5 from 8 to 7
move 1 from 3 to 1
move 2 from 8 to 5
move 2 from 8 to 1
move 1 from 4 to 1
move 6 from 2 to 7
move 3 from 5 to 3
move 1 from 2 to 5
move 7 from 7 to 9
move 3 from 3 to 5
move 1 from 2 to 5
move 2 from 3 to 2
move 6 from 1 to 7
move 10 from 7 to 3
move 1 from 2 to 3
move 6 from 9 to 8
move 1 from 2 to 4
move 2 from 6 to 1
move 5 from 1 to 9
move 8 from 5 to 8
move 2 from 1 to 6
move 6 from 3 to 4
move 1 from 5 to 3
move 4 from 9 to 6
move 1 from 1 to 4
move 2 from 9 to 2
move 5 from 6 to 1
move 11 from 6 to 7
move 1 from 2 to 8
move 6 from 7 to 5
move 10 from 8 to 4
move 2 from 3 to 9
move 3 from 3 to 5
move 4 from 7 to 9
move 2 from 1 to 3
move 10 from 5 to 8
move 6 from 6 to 1
move 2 from 6 to 8
move 2 from 9 to 5
move 4 from 9 to 6
move 7 from 4 to 8
move 5 from 6 to 1
move 4 from 8 to 2
move 2 from 5 to 6
move 5 from 4 to 5
move 1 from 7 to 5
move 2 from 3 to 6
move 1 from 3 to 8
move 4 from 6 to 1
move 4 from 2 to 3
move 5 from 5 to 1
move 2 from 3 to 2
move 2 from 3 to 2
move 20 from 8 to 2
move 5 from 4 to 8
move 1 from 4 to 3
move 8 from 2 to 1
move 1 from 5 to 6
move 5 from 2 to 3
move 1 from 6 to 5
move 5 from 3 to 2
move 1 from 3 to 7
move 6 from 8 to 5
move 13 from 2 to 9
move 7 from 9 to 8
move 1 from 7 to 8
move 5 from 8 to 3
move 2 from 2 to 5
move 2 from 8 to 4
move 27 from 1 to 5
move 1 from 2 to 3
move 5 from 3 to 1
move 22 from 5 to 7
move 1 from 8 to 5
move 1 from 3 to 2
move 7 from 1 to 3
move 2 from 3 to 7
move 2 from 2 to 4
move 5 from 9 to 1
move 5 from 3 to 9
move 3 from 1 to 5
move 3 from 1 to 6
move 3 from 6 to 3
move 4 from 4 to 2
move 8 from 5 to 3
move 8 from 7 to 4
move 14 from 7 to 4
move 1 from 1 to 7
move 6 from 9 to 6
move 7 from 5 to 3
move 14 from 3 to 6
move 2 from 2 to 1
move 4 from 3 to 7
move 6 from 7 to 6
move 1 from 7 to 6
move 1 from 5 to 1
move 2 from 1 to 5
move 3 from 5 to 7
move 8 from 6 to 5
move 5 from 5 to 1
move 1 from 7 to 3
move 1 from 3 to 8
move 22 from 4 to 7
move 7 from 6 to 3
move 4 from 3 to 2
move 3 from 1 to 3
move 17 from 7 to 6
move 1 from 8 to 1
move 2 from 2 to 4
move 3 from 7 to 2
move 2 from 2 to 9
move 1 from 1 to 8
move 2 from 3 to 1
move 6 from 6 to 8
move 2 from 9 to 2
move 4 from 5 to 1
move 5 from 8 to 9
move 1 from 7 to 3
move 4 from 3 to 4
move 1 from 7 to 4
move 4 from 9 to 7
move 5 from 7 to 9
move 1 from 7 to 3
move 2 from 2 to 8
move 5 from 4 to 2
move 21 from 6 to 8
move 2 from 3 to 8
move 23 from 8 to 6
move 1 from 2 to 6
move 2 from 9 to 8
move 22 from 6 to 7
move 2 from 9 to 3
move 2 from 3 to 7
move 2 from 1 to 6
move 1 from 2 to 5
move 3 from 1 to 3
move 6 from 7 to 4
move 5 from 8 to 5
move 1 from 3 to 8
move 1 from 9 to 3
move 6 from 4 to 8
move 1 from 5 to 3
move 6 from 2 to 8
move 15 from 7 to 5
move 1 from 7 to 1
move 14 from 5 to 8
move 1 from 4 to 9
move 5 from 1 to 7
move 3 from 6 to 2
move 4 from 5 to 6
move 1 from 4 to 8
move 4 from 3 to 1
move 2 from 9 to 2
move 7 from 7 to 1
move 7 from 2 to 7
move 9 from 8 to 6
move 7 from 7 to 1
move 12 from 6 to 8
move 25 from 8 to 6
move 3 from 8 to 1
move 28 from 6 to 2
move 15 from 2 to 3
move 1 from 5 to 4
move 3 from 2 to 7
move 6 from 2 to 9
`
