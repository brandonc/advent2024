package day04

import (
	"bufio"
	"io"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day04 struct {
	letters [][]byte
}

func Factory() solution.Solver {
	return day04{}
}

func (d *day04) load(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	d.letters = make([][]byte, 0)

	ui.Debug("Loading letters")
	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())

		ui.Debug((string)(line))
		d.letters = append(d.letters, line)
	}
}

func (d *day04) horizRight(x, y int) bool {
	return len(d.letters[y]) > x+3 && d.letters[y][x+1] == 'M' && d.letters[y][x+2] == 'A' && d.letters[y][x+3] == 'S'
}

func (d *day04) horizLeft(x, y int) bool {
	return x >= 3 && d.letters[y][x-1] == 'M' && d.letters[y][x-2] == 'A' && d.letters[y][x-3] == 'S'
}

func (d *day04) vertDown(x, y int) bool {
	return len(d.letters) > y+3 && d.letters[y+1][x] == 'M' && d.letters[y+2][x] == 'A' && d.letters[y+3][x] == 'S'
}

func (d *day04) vertUp(x, y int) bool {
	return y >= 3 && d.letters[y-1][x] == 'M' && d.letters[y-2][x] == 'A' && d.letters[y-3][x] == 'S'
}

func (d *day04) diagDownRight(x, y int) bool {
	// y+, x+
	return len(d.letters) > y+3 && len(d.letters[y]) > x+3 && d.letters[y+1][x+1] == 'M' && d.letters[y+2][x+2] == 'A' && d.letters[y+3][x+3] == 'S'
}

func (d *day04) diagUpLeft(x, y int) bool {
	// y-, x-
	return x >= 3 && y >= 3 && d.letters[y-1][x-1] == 'M' && d.letters[y-2][x-2] == 'A' && d.letters[y-3][x-3] == 'S'
}

func (d *day04) diagDownLeft(x, y int) bool {
	// y+, x-
	return len(d.letters) > y+3 && x >= 3 && d.letters[y+1][x-1] == 'M' && d.letters[y+2][x-2] == 'A' && d.letters[y+3][x-3] == 'S'
}

func (d *day04) diagUpRight(x, y int) bool {
	// y-, x+
	return y >= 3 && len(d.letters[y]) > x+3 && d.letters[y-1][x+1] == 'M' && d.letters[y-2][x+2] == 'A' && d.letters[y-3][x+3] == 'S'
}

func (d day04) Part1(reader io.Reader) int {
	d.load(reader)

	count := 0
	for y := 0; y < len(d.letters); y++ {
		for x := 0; x < len(d.letters[y]); x++ {
			if d.letters[y][x] != 'X' {
				continue
			}

			if d.vertUp(x, y) {
				ui.Debugf("Found vert up at %d, %d", y, x)
				count += 1
			}

			if d.vertDown(x, y) {
				ui.Debugf("Found vert down at %d, %d", y, x)
				count += 1
			}

			if d.horizLeft(x, y) {
				ui.Debugf("Found horiz left at %d, %d", y, x)
				count += 1
			}

			if d.horizRight(x, y) {
				ui.Debugf("Found horiz right at %d, %d", y, x)
				count += 1
			}

			if d.diagUpLeft(x, y) {
				ui.Debugf("Found diag up left at %d, %d", y, x)
				count += 1
			}

			if d.diagUpRight(x, y) {
				ui.Debugf("Found diag up right at %d, %d", y, x)
				count += 1
			}

			if d.diagDownLeft(x, y) {
				ui.Debugf("Found diag down left at %d, %d", y, x)
				count += 1
			}

			if d.diagDownRight(x, y) {
				ui.Debugf("Found diag down right at %d, %d", y, x)
				count += 1
			}
		}
	}

	return count
}

func (d *day04) wordDownRight(y, x int) [3]byte {
	word := [3]byte{0, 0, 0}
	if y < 0 || x < 0 || y+2 >= len(d.letters) || x+2 >= len(d.letters[y]) {
		return word
	}

	word[0] = d.letters[y][x]
	word[1] = d.letters[y+1][x+1]
	word[2] = d.letters[y+2][x+2]

	return word
}

func (d *day04) wordDownLeft(y, x int) [3]byte {
	word := [3]byte{0, 0, 0}
	if y < 0 || x >= len(d.letters[0]) || y+2 >= len(d.letters) || x <= 1 {
		return word
	}

	word[0] = d.letters[y][x]
	word[1] = d.letters[y+1][x-1]
	word[2] = d.letters[y+2][x-2]

	return word
}

func (d day04) Part2(reader io.Reader) int {
	d.load(reader)

	mas := [3]byte{'M', 'A', 'S'}
	sam := [3]byte{'S', 'A', 'M'}

	count := 0
	for y := 0; y < len(d.letters); y++ {
		for x := 0; x < len(d.letters[y]); x++ {
			if d.letters[y][x] != 'A' {
				continue
			}

			right := d.wordDownRight(y-1, x-1)
			left := d.wordDownLeft(y-1, x+1)

			if (left == mas || left == sam) && (right == mas || right == sam) {
				ui.Debugf("Found X-MAS at %d, %d", y, x)
				count += 1
			}
		}
	}

	return count
}
