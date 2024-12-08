package day08

import (
	"bufio"
	"io"

	"github.com/brandonc/advent2024/solutions/solution"
)

type position struct {
	x int
	y int
}

type day08 struct {
	antennas map[byte][]position
	fieldY   int
	fieldX   int
}

func Factory() solution.Solver {
	return day08{}
}

func (d *day08) load(reader io.Reader) {
	d.antennas = make(map[byte][]position)

	scanner := bufio.NewScanner(reader)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		d.fieldX = len(line)

		for x, b := range []byte(line) {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
				if _, exist := d.antennas[b]; exist {
					d.antennas[b] = append(d.antennas[b], position{y: y, x: x})
				} else {
					d.antennas[b] = []position{{y: y, x: x}}
				}
			}
		}
		y += 1
	}
	d.fieldY = y
}

func antinode(i, j position) position {
	antinode := position{}

	// Calulate antinode when i is to the left of j
	if i.x < j.x {
		// 8,8 + 9,9 = 7,7
		antinode.x = i.x - (j.x - i.x)
	} else {
		// 9,9 + 8,8 = 10,10
		antinode.x = i.x + (i.x - j.x)
	}

	// Calulate antinode when i is above j
	if i.y < j.y {
		// 8,8 + 9,9 = 7,7
		antinode.y = i.y - (j.y - i.y)
	} else {
		// 9,9 + 8,8 = 10,10
		antinode.y = i.y + (i.y - j.y)
	}

	return antinode
}

func (d day08) Part1(reader io.Reader) int {
	d.load(reader)

	locations := make(map[position]any)
	for _, antennas := range d.antennas {
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}

				an := antinode(antennas[i], antennas[j])

				if an.x < 0 || an.x >= d.fieldX || an.y < 0 || an.y >= d.fieldY {
					continue
				}

				if _, ok := locations[an]; !ok {
					locations[an] = struct{}{}
				}
			}
		}
	}

	return len(locations)
}

func (d day08) Part2(reader io.Reader) int {
	d.load(reader)

	locations := make(map[position]any)
	for _, antennas := range d.antennas {
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}

				an := antinode(antennas[i], antennas[j])
				previous := antennas[i]

				for {
					if an.x < 0 || an.x >= d.fieldX || an.y < 0 || an.y >= d.fieldY {
						break
					}

					if _, ok := locations[an]; !ok {
						locations[an] = struct{}{}
					}

					// antinode is at 7,7 based on 8,8 and 9,9... The next iteration should find the antinode at 6,6 based on 7,7 and 8,8
					next := antinode(an, previous)
					previous = an

					an = position{y: next.y, x: next.x}
				}
			}
		}

		// Add antinode positions at every antenna position as long as there are at least 2 antennas
		if len(antennas) > 1 {
			for _, antenna := range antennas {
				if _, ok := locations[antenna]; !ok {
					locations[antenna] = struct{}{}
				}
			}
		}
	}

	return len(locations)
}
