package day06

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day06 struct {
	lab      [][]byte
	posY     int
	posX     int
	initX    int
	initY    int
	facing   byte
	distinct int
}

type vector struct {
	pos position
	dir byte
}

type position struct {
	y int
	x int
}

type finished byte

var leftLab finished = 'L'
var barrier finished = 'B'

func Factory() solution.Solver {
	return day06{}
}

func (d *day06) reset() {
	for y := 0; y < len(d.lab); y++ {
		for x := 0; x < len(d.lab[y]); x++ {
			if d.lab[y][x] == 'X' {
				d.lab[y][x] = '.'
			}
		}
	}
	d.posY, d.posX = d.initY, d.initX
	d.lab[d.posY][d.posX] = 'X'
	d.facing = '^'
	d.distinct = 1
}

func (d *day06) load(reader io.Reader) {
	d.lab = make([][]byte, 0)
	d.distinct = 1

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if indexX := strings.IndexByte(line, '^'); indexX >= 0 {
			d.posX, d.posY = indexX, len(d.lab)
			d.initX, d.initY = d.posX, d.posY
		}

		d.lab = append(d.lab, []byte(line))
	}

	d.lab[d.posY][d.posX] = 'X'
	d.facing = '^'
}

func (d *day06) forward() finished {
	for {
		switch d.facing {
		case '^':
			if d.posY == 0 {
				return leftLab
			} else if d.lab[d.posY-1][d.posX] == '#' {
				return barrier
			}
			d.posY--
		case 'v':
			if d.posY == len(d.lab)-1 {
				return leftLab
			} else if d.lab[d.posY+1][d.posX] == '#' {
				return barrier
			}
			d.posY++
		case '<':
			if d.posX == 0 {
				return leftLab
			} else if d.lab[d.posY][d.posX-1] == '#' {
				return barrier
			}
			d.posX--
		case '>':
			if d.posX == len(d.lab[d.posY])-1 {
				return leftLab
			} else if d.lab[d.posY][d.posX+1] == '#' {
				return barrier
			}
			d.posX++
		}

		if d.lab[d.posY][d.posX] == '.' {
			d.lab[d.posY][d.posX] = 'X'
			d.distinct++
		}
	}
}

func (d *day06) start() bool {
	barriersSeen := make([]vector, 0)

	for {
		// Reorient if near a wall
		switch d.facing {
		case '^':
			if d.posY > 0 && d.lab[d.posY-1][d.posX] == '#' {
				d.facing = '>'
			}
		case 'v':
			if d.posY < len(d.lab)-1 && d.lab[d.posY+1][d.posX] == '#' {
				d.facing = '<'
			}
		case '<':
			if d.posX > 0 && d.lab[d.posY][d.posX-1] == '#' {
				d.facing = '^'
			}
		case '>':
			if d.posX < len(d.lab[d.posY])-1 && d.lab[d.posY][d.posX+1] == '#' {
				d.facing = 'v'
			}
		}

		if d.forward() == leftLab {
			break
		}

		for _, b := range barriersSeen {
			if b.dir == d.facing && b.pos.y == d.posY && b.pos.x == d.posX {
				// Been here in this orientation before
				return false
			}
		}
		barriersSeen = append(barriersSeen, vector{dir: d.facing, pos: position{x: d.posX, y: d.posY}})
	}
	return true
}

func (d day06) Part1(reader io.Reader) int {
	d.load(reader)

	if !d.start() {
		ui.Die(errors.New("infinite loop detected"))
	}

	return d.distinct
}

func (d day06) Part2(reader io.Reader) int {
	d.load(reader)

	possibilities := 0

	for y := 0; y < len(d.lab); y++ {
		for x := 0; x < len(d.lab[y]); x++ {
			if d.lab[y][x] == '#' {
				continue
			}

			// Exhaustively check all the '.' positions with a '#' obstacle
			d.reset()
			d.lab[y][x] = '#'

			// Find out if an obstacle is at this position will make the path infinite
			if !d.start() {
				possibilities += 1
			}

			// Reset the new obstacle
			d.lab[y][x] = '.'
		}
	}
	return possibilities
}
