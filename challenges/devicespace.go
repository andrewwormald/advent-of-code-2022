package challenges

import (
	"strconv"
	"strings"
)

type SystemAnalysis struct {
	Allocations map[string]int64
}

func DeviceSpace(input string) (int64, int64) {
	analysis := SystemAnalysis{
		Allocations: make(map[string]int64),
	}
	ls := strings.Split(input, "\n")
	var curDir string
	for _, l := range ls {
		if strings.Contains(l, "$ cd ..") {
			parts := strings.Split(curDir, "/")
			curDir = strings.Join(parts[:len(parts)-1], "/")
			if curDir == "" {
				curDir = "/"
			}
			continue
		} else if strings.Contains(l, "$ cd") {
			dirName := strings.Split(l, "$ cd")[1]
			dirName = strings.ReplaceAll(dirName, " ", "")
			curDir = strings.Join([]string{curDir, dirName}, "/")
			curDir = strings.ReplaceAll(curDir, "//", "/")
			continue
		}

		if strings.Contains(l, "$ ls") {
			continue
		}

		sections := strings.Split(l, " ")
		for _, s := range sections {
			sizNum, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				continue
			}

			analysis.Allocations[curDir] += sizNum

			parts := strings.Split(curDir, "/")
			current := parts[len(parts)-1:][0]
			var rebuilt string
			for _, v := range parts {
				if v == current {
					break
				}

				if v == "" {
					rebuilt = "/"
				} else {
					rebuilt += "/" + v
				}

				rebuilt = strings.ReplaceAll(rebuilt, "//", "/")
				analysis.Allocations[rebuilt] += sizNum
			}
		}
	}

	var total int64
	for _, size := range analysis.Allocations {
		if size <= 100000 {
			total += size
		}
	}

	smallest := total
	toDelete := 30000000 - (70000000 - analysis.Allocations["/"])
	for _, size := range analysis.Allocations {
		if size >= toDelete && size < smallest {
			smallest = size
		}
	}

	return total, smallest
}
