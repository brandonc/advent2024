package day12

import (
	"strings"
	"testing"
)

var (
	sample1 = `AAAA
BBCD
BBCC
EEEC`

	sample2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

	sample3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 140; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}

	answer2 := Factory().Part1(strings.NewReader(sample2))

	if expected := 772; answer2 != expected {
		t.Fatalf("Expected answer2 1 to be %d, got %d", expected, answer2)
	}

	answer3 := Factory().Part1(strings.NewReader(sample3))

	if expected := 1930; answer3 != expected {
		t.Fatalf("Expected answer3 1 to be %d, got %d", expected, answer3)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample1))

	if expected := 80; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}

	answer2 := Factory().Part2(strings.NewReader(sample2))

	if expected := 436; answer2 != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer2)
	}

	answer3 := Factory().Part2(strings.NewReader(sample3))

	if expected := 1206; answer3 != expected {
		t.Fatalf("Expected answer3 1 to be %d, got %d", expected, answer3)
	}
}
