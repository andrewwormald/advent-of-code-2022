package challenges

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MonkeyInTheMiddle(input string, rounds int) int {
	lines := strings.Split(input, "\n")

	t := new(troop)
	m := &monkey{
		others: make(map[int]*monkey),
	}

	for _, line := range lines {
		if line == "" {
			t.Add(m)
			m = &monkey{
				others: make(map[int]*monkey),
			}
			continue
		}

		_, num, ok := strings.Cut(line, "Monkey ")
		if ok {
			num = strings.Split(num, ":")[0]
			id, _ := strconv.Atoi(num)
			m.id = id
		}

		_, after, ok := strings.Cut(line, "Starting items: ")
		if ok {
			after = strings.ReplaceAll(after, " ", "")
			nums := strings.Split(after, ",")
			for _, s := range nums {
				worry, _ := strconv.Atoi(s)
				m.items = append(m.items, worry)
			}
		}

		_, after, ok = strings.Cut(line, "Operation: new = old ")
		if ok {
			parts := strings.Split(after, " ")
			operator := parts[0]
			val := parts[1]

			m.operation = func(item int) int {
				worry, _ := strconv.Atoi(val)
				if val == "old" {
					worry = item
				}

				switch operator {
				case "+":
					return item + worry
				case "-":
					return item - worry
				case "/":
					return int(math.Floor(float64(item) / float64(worry)))
				case "*":
					return int(math.Floor(float64(item) * float64(worry)))
				default:
					panic("dont know which operator to use")
				}

				return 0
			}
		}

		_, after, ok = strings.Cut(line, "Test: divisible by ")
		if ok {
			divisBy, _ := strconv.Atoi(after)
			m.divisBy = divisBy
			m.test = func(item int) bool {
				return item%divisBy == 0
			}
		}

		_, after, ok = strings.Cut(line, "If true: throw to monkey ")
		if ok {
			id, _ := strconv.Atoi(after)
			m.ifTrueThrowTo = id
		}

		_, after, ok = strings.Cut(line, "If false: throw to monkey ")
		if ok {
			id, _ := strconv.Atoi(after)
			m.ifFalseThrowTo = id
		}
	}

	t.Sync()

	for i := 0; i < rounds; i++ {
		t.TakeRound()
	}

	t.ListInspections()

	var mostActive int
	var secondMostActive int

	for _, inspection := range t.AllInspections() {
		if mostActive < inspection {
			secondMostActive = mostActive
			mostActive = inspection
		} else if secondMostActive < inspection {
			secondMostActive = inspection
		}
	}

	return mostActive * secondMostActive
}

type troop []*monkey

func (t *troop) Add(m *monkey) {
	temp := *t
	temp = append(temp, m)
	*t = temp
}

func (t *troop) TakeRound() {
	for _, monkey := range *t {
		monkey.TakeTurn()
	}
}

func (t *troop) AllInspections() []int {
	var ls []int
	for _, m := range *t {
		ls = append(ls, m.inspections)
	}
	return ls
}

func (t *troop) ListInspections() {
	for _, m := range *t {
		fmt.Println(
			fmt.Sprintf("Moneky %v inspected items %v times", m.id, m.inspections),
		)
	}
}

func (t *troop) Sync() {
	for _, m := range *t {
		for _, m2 := range *t {
			if m.id == m2.id {
				continue
			}

			m.others[m2.id] = m2
		}
	}
}

type monkey struct {
	id int

	items     []int
	operation func(item int) int

	test           func(item int) bool
	divisBy        int
	ifTrueThrowTo  int
	ifFalseThrowTo int

	inspections int

	others map[int]*monkey
}

func (m *monkey) TakeTurn() {
	for _, item := range m.items {
		m.inspections++
		newVal := m.operation(item)

		// Part 1
		//newVal = int(math.Floor(float64(newVal) / 3))

		// Part 2
		start := m.divisBy
		for _, m2 := range m.others {
			start = start * m2.divisBy
		}
		newVal = newVal % start

		var throwTo int
		if m.test(newVal) {
			throwTo = m.ifTrueThrowTo
		} else {
			throwTo = m.ifFalseThrowTo
		}

		m.ThrowTo(throwTo, newVal)
	}

	m.items = []int{}
}

func (m *monkey) ThrowTo(id int, item int) {
	m.others[id].items = append(m.others[id].items, item)
}
