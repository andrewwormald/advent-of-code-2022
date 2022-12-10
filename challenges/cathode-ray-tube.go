package challenges

import (
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	cyc    int
	x      int
	signal int
}

func (c *cpu) cycle() {
	if c.cyc%40 == 0 {
		fmt.Print("\n")
	} else if c.cyc%40 == c.x {
		fmt.Print("#")
	} else if c.cyc%40 == c.x-1 {
		fmt.Print("#")
	} else if c.cyc%40 == c.x+1 {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}

	c.cyc++
	if (c.cyc-20)%40 != 0 {
		return
	}

	c.signal += c.x * c.cyc
}

func Cathode(input string) int {
	lines := strings.Split(input, "\n")

	cpu := &cpu{
		x: 1,
	}

	for _, ln := range lines {
		if ln == "" {
			continue
		}

		if ln == "noop" {
			cpu.cycle()
		} else {
			cpu.cycle()
			cpu.cycle()

			_, val, _ := strings.Cut(ln, " ")
			num, _ := strconv.Atoi(val)
			cpu.x += num
		}
	}

	return cpu.signal
}
