package day07

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day07 struct {
	tests               []testCase
	allowConcatOperator bool
}

type testCase struct {
	answer   int
	operands []int
}

func Factory() solution.Solver {
	return day07{}
}

func (d *day07) load(reader io.Reader) {
	result := make([]testCase, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		records := strings.Split(line, ": ")

		if len(records) != 2 {
			ui.Die(fmt.Errorf("invalid test case %q", line))
		}

		operands := strings.Split(records[1], " ")

		answer, err := strconv.Atoi(records[0])
		if err != nil {
			ui.Die(fmt.Errorf("invalid test case %q", line))
		}

		test := testCase{
			answer:   answer,
			operands: make([]int, len(operands)),
		}

		for i, operand := range operands {
			op, err := strconv.Atoi(operand)
			if err != nil {
				ui.Die(fmt.Errorf("invalid test case %q", line))
			}
			test.operands[i] = op
		}

		result = append(result, test)
	}

	d.tests = result
}

func (d day07) anyPossibilitiesTrue(operands []int, answer int) bool {
	if len(operands) == 0 {
		ui.Die(fmt.Errorf("no operands"))
	}

	if len(operands) == 1 {
		return operands[0] == answer
	}

	// If we've already blown it, no more operations will help
	if operands[0] > answer {
		return false
	}

	firstPossibility := make([]int, len(operands)-1)
	secondPossibility := make([]int, len(operands)-1)
	thirdPossibility := make([]int, len(operands)-1)

	firstPossibility[0] = operands[0] * operands[1]
	secondPossibility[0] = operands[0] + operands[1]

	third, err := strconv.Atoi(fmt.Sprintf("%d%d", operands[0], operands[1]))
	if err != nil {
		ui.Die(fmt.Errorf("invalid test case %q", operands))
	}

	thirdPossibility[0] = third

	for i := 2; i < len(operands); i++ {
		firstPossibility[i-1] = operands[i]
		secondPossibility[i-1] = operands[i]
		thirdPossibility[i-1] = operands[i]
	}

	if d.anyPossibilitiesTrue(firstPossibility, answer) || d.anyPossibilitiesTrue(secondPossibility, answer) {
		return true
	} else if d.allowConcatOperator {
		return d.anyPossibilitiesTrue(thirdPossibility, answer)
	}

	return false
}

func (d day07) Part1(reader io.Reader) int {
	d.load(reader)

	total := 0
	for _, test := range d.tests {
		if d.anyPossibilitiesTrue(test.operands, test.answer) {
			total += test.answer
		}
	}

	return total
}

func (d day07) Part2(reader io.Reader) int {
	d.load(reader)
	d.allowConcatOperator = true

	total := 0
	for _, test := range d.tests {
		if d.anyPossibilitiesTrue(test.operands, test.answer) {
			total += test.answer
		}
	}

	return total
}
