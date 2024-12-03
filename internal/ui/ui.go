package ui

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/colorstring"
)

func Die(err error) {
	if err != nil {
		colorstring.Printf("[red]An unexpected error occurred:\n%s[reset]\n", err)
		os.Exit(1)
	}
}

func Assert(expr bool, description string) {
	if !expr {
		Die(errors.New(description))
		os.Exit(1)
	}
}

func Debug(message string) {
	if os.Getenv("LOG_LEVEL") != "debug" {
		return
	}
	colorstring.Printf("[dark_gray][DEBUG] %s\n", message)
}

func Debugf(message string, a ...any) {
	Debug(fmt.Sprintf(message, a...))
}

func rightAlign(v string, other ...string) string {
	maxOther := 0
	for _, o := range other {
		if len(o) > maxOther {
			maxOther = len(o)
		}
	}
	if len(v) > maxOther {
		return v
	} else {
		return fmt.Sprintf("%s%s", strings.Repeat(" ", maxOther-len(v)), v)
	}
}

func humanizeDuration(d time.Duration) string {
	units := []struct {
		unit   string
		amount int
	}{
		{"s", int(d.Seconds())},
		{"ms", int(d.Milliseconds())},
		{"μs", int(d.Microseconds())},
		{"ns", int(d.Nanoseconds())},
	}

	for _, u := range units {
		if u.amount == 0 {
			continue
		}

		return fmt.Sprintf("%d%s", u.amount, u.unit)
	}

	return "0 ns"
}

func Answer(part1 func() int, part2 func() int) {
	startTimePart1 := time.Now()
	timePart1 := time.Since(startTimePart1)

	startTimePart2 := time.Now()
	timePart2 := time.Since(startTimePart2)

	colorstring.Printf("[cyan]Part 1\n[white]%d\n[dark_gray]%s\n\n", part1(), humanizeDuration(timePart1))
	colorstring.Printf("[cyan]Part 2\n[white]%d\n[dark_gray]%s\n\n", part2(), humanizeDuration(timePart2))

	colorstring.Printf("[dark_gray]total duration %s\n", humanizeDuration(timePart1+timePart2))
}
