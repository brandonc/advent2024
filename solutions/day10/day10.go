package day10

import (
	"bufio"
	"io"

	"github.com/brandonc/advent2024/solutions/solution"
)

type day10 struct {
	grid       [][]int
	trailheads []position
	peaks      []position
}

type position struct {
	y, x int
}

func (d *day10) load(reader io.Reader) {
	d.grid = make([][]int, 0)
	d.trailheads = make([]position, 0)

	scanner := bufio.NewScanner(reader)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for x, c := range line {
			if c == '0' {
				d.trailheads = append(d.trailheads, position{y, x})
			}
			row[x] = int(c - '0')
		}
		d.grid = append(d.grid, row)
		y++
	}
}

func Factory() solution.Solver {
	return day10{}
}

func (d *day10) score(step, y, x int, reached map[position]any) int {
	if step == 9 {
		if _, ok := reached[position{y, x}]; !ok {
			reached[position{y, x}] = struct{}{}
			return 1
		} else {
			return 0
		}
	}

	score := 0
	// up
	if y > 0 && d.grid[y-1][x] == step+1 {
		score += d.score(step+1, y-1, x, reached)
	}

	// down
	if y < len(d.grid)-1 && d.grid[y+1][x] == step+1 {
		score += d.score(step+1, y+1, x, reached)
	}

	// left
	if x > 0 && d.grid[y][x-1] == step+1 {
		score += d.score(step+1, y, x-1, reached)
	}

	// right
	if x < len(d.grid[0])-1 && d.grid[y][x+1] == step+1 {
		score += d.score(step+1, y, x+1, reached)
	}

	return score
}

func (d *day10) rate(step, y, x int) int {
	if step == 9 {
		return 1
	}

	score := 0
	// up
	if y > 0 && d.grid[y-1][x] == step+1 {
		score += d.rate(step+1, y-1, x)
	}

	// down
	if y < len(d.grid)-1 && d.grid[y+1][x] == step+1 {
		score += d.rate(step+1, y+1, x)
	}

	// left
	if x > 0 && d.grid[y][x-1] == step+1 {
		score += d.rate(step+1, y, x-1)
	}

	// right
	if x < len(d.grid[0])-1 && d.grid[y][x+1] == step+1 {
		score += d.rate(step+1, y, x+1)
	}

	return score
}

func (d day10) Part1(reader io.Reader) int {
	d.load(reader)

	score := 0
	for _, th := range d.trailheads {
		score += d.score(0, th.y, th.x, make(map[position]any))
	}

	return score
}

func (d day10) Part2(reader io.Reader) int {
	d.load(reader)

	rating := 0
	for _, th := range d.trailheads {
		rating += d.rate(0, th.y, th.x)
	}

	return rating
}
