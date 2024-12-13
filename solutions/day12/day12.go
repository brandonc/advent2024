package day12

import (
	"bufio"
	"io"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type position struct {
	y, x int
}

type plot struct {
	pos                            position
	fenceN, fenceE, fenceS, fenceW bool
	plant                          byte
}

type day12 struct {
	field [][]plot
}

func Factory() solution.Solver {
	return day12{}
}

func (d *day12) load(reader io.Reader) {
	d.field = make([][]plot, 0)

	y := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]plot, len(line))

		for i := 0; i < len(line); i++ {
			pos := position{y: y, x: i}
			row[i] = plot{pos: pos, plant: line[i], fenceN: true, fenceE: true, fenceS: true, fenceW: true}
		}

		d.field = append(d.field, row)
		y++
	}
}

func (d *day12) connectPlots() {
	for y := 0; y < len(d.field); y++ {
		for x := 0; x < len(d.field[y]); x++ {
			current := d.field[y][x].plant
			if y > 0 {
				if d.field[y-1][x].plant == current {
					d.field[y][x].fenceN = false
					d.field[y-1][x].fenceS = false
				}
			}

			if y < len(d.field)-1 {
				if d.field[y+1][x].plant == current {
					d.field[y][x].fenceS = false
					d.field[y+1][x].fenceN = false
				}
			}

			if x > 0 {
				if d.field[y][x-1].plant == current {
					d.field[y][x].fenceW = false
					d.field[y][x-1].fenceE = false
				}
			}

			if x < len(d.field[y])-1 {
				if d.field[y][x+1].plant == current {
					d.field[y][x].fenceE = false
					d.field[y][x+1].fenceW = false
				}
			}
		}
	}
}

func (d *day12) sameN(p position, plant byte) bool {
	if p.y == 0 {
		return false
	}

	return d.field[p.y-1][p.x].plant == plant
}

func (d *day12) sameS(p position, plant byte) bool {
	if p.y == len(d.field)-1 {
		return false
	}

	return d.field[p.y+1][p.x].plant == plant
}

func (d *day12) sameW(p position, plant byte) bool {
	if p.x == 0 {
		return false
	}

	return d.field[p.y][p.x-1].plant == plant
}

func (d *day12) sameE(p position, plant byte) bool {
	if p.x == len(d.field[p.x])-1 {
		return false
	}

	return d.field[p.y][p.x+1].plant == plant
}

func (d *day12) sameSW(p position, plant byte) bool {
	if p.y == len(d.field)-1 || p.x == 0 {
		return false
	}

	return d.field[p.y+1][p.x-1].plant == plant
}

func (d *day12) sameSE(p position, plant byte) bool {
	if p.y == len(d.field)-1 || p.x == len(d.field[p.y])-1 {
		return false
	}

	return d.field[p.y+1][p.x+1].plant == plant
}

func (d *day12) sameNE(p position, plant byte) bool {
	if p.y == 0 || p.x == len(d.field[p.y])-1 {
		return false
	}

	return d.field[p.y-1][p.x+1].plant == plant
}

func (d *day12) sameNW(p position, plant byte) bool {
	if p.y == 0 || p.x == 0 {
		return false
	}

	return d.field[p.y-1][p.x-1].plant == plant
}

func (d *day12) measure(visited map[position]any, p position) (int, int, int) {
	if _, ok := visited[p]; ok {
		return 0, 0, 0
	}

	area := 1
	perimeter := 0
	corners := 0
	visited[p] = struct{}{}
	plot := d.field[p.y][p.x]

	// Outside top/left corner
	if !d.sameN(p, plot.plant) && !d.sameW(p, plot.plant) {
		ui.Debugf("Found outside top/left corner for %c at (%d, %d)", plot.plant, p.y, p.x)
		corners++
	}

	// Outside top/right corner
	if !d.sameN(p, plot.plant) && !d.sameE(p, plot.plant) {
		ui.Debugf("Found outside top/right corner for %c at (%d, %d)", plot.plant, p.y, p.x)
		corners++
	}

	// Outside bottom/left corner
	if !d.sameS(p, plot.plant) && !d.sameW(p, plot.plant) {
		ui.Debugf("Found outside bottom/left corner for %c at (%d, %d)", plot.plant, p.y, p.x)
		corners++
	}

	// Outside bottom/right corner
	if !d.sameS(p, plot.plant) && !d.sameE(p, plot.plant) {
		ui.Debugf("Found outside bottom/right corner for %c at (%d, %d)", plot.plant, p.y, p.x)
		corners++
	}

	if d.sameN(p, plot.plant) && d.sameE(p, plot.plant) && !d.sameNE(p, plot.plant) {
		ui.Debugf("Found inside corner NE for %c at (%d, %d)", plot.plant, p.y, p.x)
		corners++
	}

	if d.sameN(p, plot.plant) && d.sameW(p, plot.plant) && !d.sameNW(p, plot.plant) {
		ui.Debugf("Found inside corner NW for %c at (%d, %d)", plot.plant, p.y, p.x)
		corners++
	}

	if d.sameS(p, plot.plant) && d.sameE(p, plot.plant) && !d.sameSE(p, plot.plant) {
		ui.Debugf("Found inside corner SE for %c at (%d, %d)", plot.plant, p.y, p.x)
		corners++
	}

	if d.sameS(p, plot.plant) && d.sameW(p, plot.plant) && !d.sameSW(p, plot.plant) {
		ui.Debugf("Found inside corner SW for %c at (%d, %d)", plot.plant, p.y, p.x)
		corners++
	}

	if !plot.fenceN {
		areaN, perimeterN, cornersN := d.measure(visited, position{y: p.y - 1, x: p.x})
		area += areaN
		perimeter += perimeterN
		corners += cornersN
	} else {
		perimeter++
	}

	if !plot.fenceE {
		areaE, perimeterE, cornersE := d.measure(visited, position{y: p.y, x: p.x + 1})
		area += areaE
		perimeter += perimeterE
		corners += cornersE
	} else {
		perimeter++
	}

	if !plot.fenceS {
		areaS, perimeterS, cornersS := d.measure(visited, position{y: p.y + 1, x: p.x})
		area += areaS
		perimeter += perimeterS
		corners += cornersS
	} else {
		perimeter++
	}

	if !plot.fenceW {
		areaW, perimeterW, cornersW := d.measure(visited, position{y: p.y, x: p.x - 1})
		area += areaW
		perimeter += perimeterW
		corners += cornersW
	} else {
		perimeter++
	}

	return area, perimeter, corners
}

func (d day12) Part1(reader io.Reader) int {
	d.load(reader)
	d.connectPlots()

	visited := make(map[position]any)
	cost := 0
	for y := 0; y < len(d.field); y++ {
		for x := 0; x < len(d.field[y]); x++ {
			curPos := position{y: y, x: x}

			area, perimeter, corners := d.measure(visited, curPos)
			if area > 0 || perimeter > 0 {
				ui.Debugf("Total area for %c is %d", d.field[y][x].plant, area)
				ui.Debugf("Total perimeter for %c is %d", d.field[y][x].plant, perimeter)
				ui.Debugf("Total corners for %c is %d", d.field[y][x].plant, corners)
			}

			cost += area * perimeter
		}
	}

	return cost
}

func (d day12) Part2(reader io.Reader) int {
	d.load(reader)
	d.connectPlots()

	visited := make(map[position]any)
	cost := 0
	for y := 0; y < len(d.field); y++ {
		for x := 0; x < len(d.field[y]); x++ {
			curPos := position{y: y, x: x}

			area, perimeter, corners := d.measure(visited, curPos)
			if area > 0 || perimeter > 0 || corners > 0 {
				ui.Debugf("Total area for %c is %d", d.field[y][x].plant, area)
				ui.Debugf("Total perimeter for %c is %d", d.field[y][x].plant, perimeter)
				ui.Debugf("Total corners for %c is %d", d.field[y][x].plant, corners)
			}

			cost += area * corners
		}
	}

	return cost
}
