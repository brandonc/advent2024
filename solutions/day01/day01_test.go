package day01

import (
	"strings"
	"testing"
)

var (
	sample1 = `3   4
4   3
2   5
1   3
3   9
3   3`

	sample2 = `3   4
4   3
2   5
1   3
3   9
3   3`
)

func TestPart1(t *testing.T) {
	answer := Factory().Part1(strings.NewReader(sample1))

	if expected := 11; answer != expected {
		t.Fatalf("Expected answer 1 to be %d, got %d", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := Factory().Part2(strings.NewReader(sample2))

	if expected := 31; answer != expected {
		t.Fatalf("Expected answer 2 to be %d, got %d", expected, answer)
	}
}
