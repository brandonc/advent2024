package day11

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day11 struct{}

type stoneCounter map[string]int

func Factory() solution.Solver {
	return day11{}
}

func (d *day11) load(reader io.Reader) stoneCounter {
	data, err := io.ReadAll(reader)
	if err != nil {
		ui.Die(err)
	}

	result := make(stoneCounter)

	for _, n := range strings.Split(strings.TrimSpace(string(data)), " ") {
		result[n] += 1
	}

	return result
}

func (d *day11) blink(c stoneCounter) stoneCounter {
	newStones := make(stoneCounter)

	for number, qty := range c {
		if number == "0" {
			newStones["1"] += qty
		} else if len(number)%2 == 0 {
			newLeft := number[:len(number)/2]
			newRight := number[len(number)/2:]

			for len(newRight) > 0 && newRight[0] == '0' {
				newRight = newRight[1:]
			}

			if newRight == "" {
				newRight = "0"
			}

			newStones[newLeft] += qty
			newStones[newRight] += qty
		} else {
			num, err := strconv.Atoi(number)
			if err != nil {
				ui.Die(err)
			}
			newStones[fmt.Sprintf("%d", num*2024)] += qty
		}
	}

	return newStones
}

func (d day11) Part1(reader io.Reader) int {
	c := d.load(reader)

	for i := 0; i < 25; i++ {
		c = d.blink(c)
	}

	count := 0
	for _, qty := range c {
		count += qty
	}
	return count
}

func (d day11) Part2(reader io.Reader) int {
	c := d.load(reader)

	for i := 0; i < 75; i++ {
		c = d.blink(c)
	}

	count := 0
	for _, qty := range c {
		count += qty
	}
	return count
}
