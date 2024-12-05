package day05

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2024/internal/ds"
	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day05 struct {
	afters  map[int]ds.IntSet
	befores map[int]ds.IntSet

	updates [][]int
}

func Factory() solution.Solver {
	return day05{}
}

func (d *day05) load(reader io.Reader) {
	d.afters = make(map[int]ds.IntSet)
	d.befores = make(map[int]ds.IntSet)
	d.updates = make([][]int, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			ui.Die(fmt.Errorf("invalid rule: %q", line))
		}

		before, err := strconv.Atoi(parts[0])
		if err != nil {
			ui.Die(err)
		}
		after, err := strconv.Atoi(parts[1])
		if err != nil {
			ui.Die(err)
		}

		if _, ok := d.afters[before]; !ok {
			d.afters[before] = ds.IntSet{}
		}
		d.afters[before].Add(after)

		if _, ok := d.befores[after]; !ok {
			d.befores[after] = ds.IntSet{}
		}
		d.befores[after].Add(before)
	}

	for scanner.Scan() {
		line := scanner.Text()
		updateRaw := strings.Split(line, ",")
		update := make([]int, len(updateRaw))

		for i, v := range updateRaw {
			u, err := strconv.Atoi(v)
			if err != nil {
				ui.Die(err)
			}
			update[i] = u
		}

		d.updates = append(d.updates, update)
	}
}

func (d day05) validPair(x, y int) bool {
	return d.afters[x].Exists(y) && d.befores[y].Exists(x)
}

func (d day05) updateValid(update []int) bool {
	for a, b := 0, 1; a < len(update)-1; a, b = a+1, b+1 {
		if !d.validPair(update[a], update[b]) {
			return false
		}
	}

	return true
}

func (d day05) fix(update []int) []int {
	for !d.updateValid(update) {
		ui.Debugf("%v is not valid", update)

		// Swap the first two elements that are out of order until the update is valid
	swap:
		for a, b := 0, 1; a < len(update)-1; a, b = a+1, b+1 {
			if !d.validPair(update[a], update[b]) {
				ui.Debugf("Swapping %d and %d", update[a], update[b])
				update[a], update[b] = update[b], update[a]
				break swap
			}
		}
	}

	ui.Debugf("%v is valid", update)
	return update
}

func (d day05) Part1(reader io.Reader) int {
	d.load(reader)

	total := 0
	for _, update := range d.updates {
		if d.updateValid(update) {
			total += update[len(update)/2]
		}
	}

	return total
}

func (d day05) Part2(reader io.Reader) int {
	d.load(reader)

	total := 0
	for _, update := range d.updates {
		if !d.updateValid(update) {
			u := d.fix(update)
			total += update[len(u)/2]
		}
	}

	return total
}
