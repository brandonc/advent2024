package day02

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2024/solutions/solution"
)

type day02 struct {
	reports [][]int
}

func Factory() solution.Solver {
	return day02{}
}

func (d day02) reportIsSafe(report []int) bool {
	result := true
	levelsIncreasing := false
	levelsDecreasing := false

	for i, level := range report {
		if i == 0 {
			continue
		} else if i == 1 {
			levelsIncreasing = level > report[i-1]
			levelsDecreasing = level < report[i-1]

			if !levelsIncreasing && !levelsDecreasing {
				result = false
				break
			}
		}

		// Increasing within threshold
		if levelsIncreasing && level > report[i-1] && level-report[i-1] >= 1 && level-report[i-1] <= 3 {
			continue
		}

		// Decreasing within threshold
		if levelsDecreasing && level < report[i-1] && report[i-1]-level >= 1 && report[i-1]-level <= 3 {
			continue
		}

		result = false
		break
	}

	return result
}

func (d *day02) load(reader io.Reader) {
	d.reports = make([][]int, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		levelsRaw := strings.Split(line, " ")
		levels := make([]int, len(levelsRaw))

		for i, levelRaw := range levelsRaw {
			level, _ := strconv.Atoi(levelRaw)
			levels[i] = level
		}

		d.reports = append(d.reports, levels)
	}
}

func (d day02) Part1(reader io.Reader) int {
	d.load(reader)

	safe := 0
	for _, report := range d.reports {
		if d.reportIsSafe(report) {
			safe += 1
		}
	}

	return safe
}

func copyOmitIndex(report []int, omitIndex int) []int {
	newReport := make([]int, len(report)-1)

	ci := 0
	for i, level := range report {
		if i == omitIndex {
			continue
		}

		newReport[ci] = level
		ci += 1
	}

	return newReport
}

func (d day02) Part2(reader io.Reader) int {
	d.load(reader)

	safe := 0
	for _, report := range d.reports {
		if d.reportIsSafe(report) {
			safe += 1
		} else {
			// Try removing one level at a time to see if the report is safe
			for i := 0; i < len(report); i++ {
				newReport := copyOmitIndex(report, i)
				if d.reportIsSafe(newReport) {
					safe += 1
					break
				}
			}
		}
	}

	return safe
}
