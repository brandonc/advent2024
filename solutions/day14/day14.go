package day14

import (
	"bufio"
	"fmt"
	"io"

	"github.com/brandonc/advent2024/internal/maths"
	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type coord struct {
	x, y int
}

type robot struct {
	y, x   int
	vy, vx int
}

type day14 struct {
	robots         []robot
	fieldX, fieldY int
}

func Factory() solution.Solver {
	return day14{}
}

func (d *day14) load(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	d.robots = make([]robot, 0)
	for scanner.Scan() {
		line := scanner.Text()

		var y, x, vy, vx int
		n, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		if err != nil {
			ui.Die(err)
		}
		if n != 4 {
			ui.Die(fmt.Errorf("expected 4 values, got %d", n))
		}

		d.robots = append(d.robots, robot{y, x, vy, vx})
	}
}

func (d *day14) move(r *robot) {
	r.x += r.vx
	r.y += r.vy

	if r.y < 0 || r.y >= d.fieldY {
		r.y = maths.Mod(r.y, d.fieldY)
	}

	if r.x < 0 || r.x >= d.fieldX {
		r.x = maths.Mod(r.x, d.fieldX)
	}
}

func (d *day14) safetyScore() int {
	var q1, q2, q3, q4 int
	for r := 0; r < len(d.robots); r++ {
		if d.robots[r].y < d.fieldY/2 {
			if d.robots[r].x < d.fieldX/2 {
				q1++
			} else if d.robots[r].x > d.fieldX/2 {
				q2++
			}
		} else if d.robots[r].y > d.fieldY/2 {
			if d.robots[r].x < d.fieldX/2 {
				q3++
			} else if d.robots[r].x > d.fieldX/2 {
				q4++
			}
		}
	}

	return q1 * q2 * q3 * q4
}

func (d day14) Part1(reader io.Reader) int {
	d.load(reader)

	d.fieldX = 101
	d.fieldY = 103

	for i := 0; i < 100; i++ {
		for r := 0; r < len(d.robots); r++ {
			d.move(&d.robots[r])
		}
	}

	return d.safetyScore()
}

func (d *day14) print() {
	robotMap := make(map[coord]any)
	for r := 0; r < len(d.robots); r++ {
		robotMap[coord{d.robots[r].x, d.robots[r].y}] = struct{}{}
	}

	for y := 0; y < 103; y++ {
		for x := 0; x < 101; x++ {
			if _, ok := robotMap[coord{x, y}]; ok {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (d day14) Part2(reader io.Reader) int {
	d.load(reader)

	d.fieldX = 101
	d.fieldY = 103
	i := 1
	for {
		unique := make(map[coord]any)
		for r := 0; r < len(d.robots); r++ {
			d.move(&d.robots[r])
			unique[coord{d.robots[r].x, d.robots[r].y}] = struct{}{}
		}

		if len(unique) == len(d.robots) {
			d.print()
			break
		}

		i++
	}

	return i
}
