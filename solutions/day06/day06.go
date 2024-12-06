package day06

import (
	"bufio"
	"io"
	"strings"

	"github.com/brandonc/advent2024/solutions/solution"
)

type day06 struct {
	lab      [][]byte
	posY     int
	posX     int
	facing   byte
	distinct int
	finished bool
	posOptX  int
	posOptY  int
}

func Factory() solution.Solver {
	return day06{}
}

func (d *day06) load(reader io.Reader) {
	d.lab = make([][]byte, 0)
	d.finished = false
	d.distinct = 1

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if indexX := strings.IndexByte(line, '^'); indexX >= 0 {
			d.posX = indexX
			d.posY = len(d.lab)
		}

		d.lab = append(d.lab, []byte(line))
	}

	d.lab[d.posY][d.posX] = 'X'
	d.facing = '^'
}

func (d *day06) forwardUntilBarrier() {
	for {
		switch d.facing {
		case '^':
			if d.posY == 0 {
				d.finished = true
				return
			} else if d.lab[d.posY-1][d.posX] == '#' {
				return
			}
			d.posY--
		case 'v':
			if d.posY == len(d.lab)-1 {
				d.finished = true
				return
			} else if d.lab[d.posY+1][d.posX] == '#' {
				return
			}
			d.posY++
		case '<':
			if d.posX == 0 {
				d.finished = true
				return
			} else if d.lab[d.posY][d.posX-1] == '#' {
				return
			}
			d.posX--
		case '>':
			if d.posX == len(d.lab[d.posY])-1 {
				d.finished = true
				return
			} else if d.lab[d.posY][d.posX+1] == '#' {
				return
			}
			d.posX++
		}

		if d.lab[d.posY][d.posX] == '.' {
			d.lab[d.posY][d.posX] = 'X'
			d.distinct++
		}
	}
}

func (d day06) Part1(reader io.Reader) int {
	d.load(reader)

	for {
		// Reorient if near a wall
		switch d.facing {
		case '^':
			if d.posY > 0 && d.lab[d.posY-1][d.posX] == '#' {
				d.facing = '>'
			}
			d.forwardUntilBarrier()
		case 'v':
			if d.posY < len(d.lab)-1 && d.lab[d.posY+1][d.posX] == '#' {
				d.facing = '<'
			}
			d.forwardUntilBarrier()
		case '<':
			if d.posX > 0 && d.lab[d.posY][d.posX-1] == '#' {
				d.facing = '^'
			}
			d.forwardUntilBarrier()
		case '>':
			if d.posX < len(d.lab[d.posY])-1 && d.lab[d.posY][d.posX+1] == '#' {
				d.facing = 'v'
			}
			d.forwardUntilBarrier()
		}

		if d.finished {
			break
		}
	}

	return d.distinct
}

func (d day06) Part2(reader io.Reader) int {
	d.load(reader)
	d.posOptX = 0
	d.posOptY = 0

	for y := 0; y < len(d.lab); y++ {
		for x := 0; x < len(d.lab[y]); x++ {
			d.posOptX = x
			d.posOptY = y

			// Find out if an obstacle is at this position will make the path infinite

		}
	}
	return 0
}
