package day11

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day11 struct {
	root *stone
}

type stone struct {
	number string
	next   *stone
}

func (s *stone) split() {
	if len(s.number)%2 != 0 {
		ui.Die(errors.New("expected an even number of digits"))
	}

	newLeft := s.number[:len(s.number)/2]
	newRight := s.number[len(s.number)/2:]

	ui.Debugf("Splitting %s into %s and %s\n", s.number, newLeft, newRight)

	for len(newRight) > 0 && newRight[0] == '0' {
		newRight = newRight[1:]
	}

	if len(newRight) == 0 {
		newRight = "0"
	}

	s.number = newLeft
	oldNext := s.next

	s.next = &stone{number: newRight, next: oldNext}
}

func (d *day11) load(reader io.Reader) {
	d.root = &stone{}
	current := d.root

	data, err := io.ReadAll(reader)
	if err != nil {
		ui.Die(err)
	}

	for _, n := range strings.Split(strings.TrimSpace(string(data)), " ") {
		s := &stone{number: n}
		current.next = s
		current = s
	}
}

func (d *day11) applyRules() {
	current := d.root.next
	for {
		if current == nil {
			break
		}
		if current.number == "0" {
			current.number = "1"
		} else if len(current.number)%2 == 0 {
			current.split()
			current = current.next
		} else {
			num, err := strconv.Atoi(current.number)
			if err != nil {
				ui.Die(err)
			}
			current.number = fmt.Sprintf("%d", num*2024)
		}
		current = current.next
	}
}

func Factory() solution.Solver {
	return day11{}
}

func (d day11) Part1(reader io.Reader) int {
	d.load(reader)

	for i := 0; i < 25; i++ {
		d.applyRules()
	}

	count := 0
	current := d.root.next
	for {
		if current == nil {
			break
		}
		count += 1
		current = current.next
	}
	return count
}

func (d day11) Part2(reader io.Reader) int {
	return 0
}
