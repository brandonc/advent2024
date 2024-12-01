package day01

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/brandonc/advent2024/internal/maths"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day01 struct {
	list1 []int
	list2 []int
}

func Factory() solution.Solver {
	return day01{}
}

func (d *day01) load(reader io.Reader) {
	d.list1 = make([]int, 0)
	d.list2 = make([]int, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		if len(parts) != 2 {
			panic(fmt.Sprintf("Expected two parts in line %q", line))
		}

		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		d.list1 = append(d.list1, num1)
		d.list2 = append(d.list2, num2)
	}

	if len(d.list1) != len(d.list2) {
		panic("Lists are not the same length")
	}
}

func (d day01) Part1(reader io.Reader) int {
	d.load(reader)

	slices.Sort(d.list1)
	slices.Sort(d.list2)

	total := 0
	for i := 0; i < len(d.list1); i++ {
		total += maths.AbsInt(d.list1[i] - d.list2[i])
	}

	return total
}

func (d day01) Part2(reader io.Reader) int {
	d.load(reader)

	slices.Sort(d.list1)
	slices.Sort(d.list2)

	total := 0
	for i := 0; i < len(d.list1); i++ {
		found := 0
		for j := 0; j < len(d.list2); j++ {
			if d.list1[i] == d.list2[j] {
				found += 1
			} else if d.list2[j] > d.list1[i] {
				break
			}
		}

		total += d.list1[i] * found
	}

	return total
}
