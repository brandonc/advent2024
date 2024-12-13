package day13

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day13 struct {
	prizes []prize
}

type coord struct {
	x, y int
}

type prize struct {
	buttonA coord
	buttonB coord
	prize   coord
}

func Factory() solution.Solver {
	return day13{}
}

func parseCoord(format, line string) coord {
	var x, y int
	if n, err := fmt.Sscanf(line, format, &x, &y); n != 2 || err != nil {
		ui.Die(err)
	}
	return coord{x: x, y: y}
}

func (d *day13) load(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	cur := prize{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Button A: ") {
			cur.buttonA = parseCoord("Button A: X+%d, Y+%d", line)
		} else if strings.HasPrefix(line, "Button B: ") {
			cur.buttonB = parseCoord("Button B: X+%d, Y+%d", line)
		} else if strings.HasPrefix(line, "Prize: ") {
			cur.prize = parseCoord("Prize: X=%d, Y=%d", line)
			d.prizes = append(d.prizes, cur)
		}
	}
}

// idk linear algebra lol copped from reddit
func (p prize) solve() int {
	D, Dx, Dy := p.buttonA.x*p.buttonB.y-p.buttonB.x*p.buttonA.y, p.prize.x*p.buttonB.y-p.buttonB.x*p.prize.y, p.buttonA.x*p.prize.y-p.prize.x*p.buttonA.y
	if D != 0 && Dx == (Dx/D)*D && Dy == (Dy/D)*D {
		return (Dx/D)*3 + (Dy / D)
	}
	return 0
}

func (d day13) Part1(reader io.Reader) int {
	d.load(reader)

	total := 0
	for _, prize := range d.prizes {
		total += prize.solve()
	}

	return total
}

func (d day13) Part2(reader io.Reader) int {
	d.load(reader)

	total := 0
	for _, p := range d.prizes {
		p.prize.x += 10000000000000
		p.prize.y += 10000000000000

		total += p.solve()
	}

	return total
}
