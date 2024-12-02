package day02

import (
	"strings"
	"testing"
)

var (
	sample1 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	sample2 = sample1
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 2; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample2))

	if expected := 4; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
