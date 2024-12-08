package day08

import (
	"strings"
	"testing"
)

var (
	sample1 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	sample2 = sample1
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 14; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample2))

	if expected := 34; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
