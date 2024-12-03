package day03

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day03 struct{}

var (
	TypeInteger  = "INTEGER"
	TypeEOF      = "EOF"
	TypeDont     = "DONT"
	TypeDo       = "DO"
	TypeMultiply = "MUL"
	TypeUnknown  = "UNKNOWN"
)

func Factory() solution.Solver {
	return day03{}
}

type Token struct {
	Type  string
	Value string
}

func peek(r *bufio.Reader) byte {
	b, err := r.Peek(1)
	if err != nil || len(b) == 0 {
		return 0
	}
	return b[0]
}

func peekN(r *bufio.Reader, n int) string {
	b, err := r.Peek(n)
	if err != nil || len(b) == 0 {
		return ""
	}
	return string(b)
}

func next(r *bufio.Reader) Token {
	b := peek(r)
	if b == 0 {
		return Token{Type: TypeEOF, Value: ""}
	}

	if b >= '0' && b <= '9' {
		result := Token{Type: TypeInteger}
		buf := strings.Builder{}
		c := b
		for {
			buf.WriteByte(c)

			r.Discard(1)
			c = peek(r)
			if peek(r) < '0' || peek(r) > '9' {
				result.Value = buf.String()
				break
			}
		}
		return result
	}

	if b == 'd' {
		if peekN(r, 7) == "don't()" {
			r.Discard(7)
			return Token{Type: TypeDont, Value: "don't()"}
		}
		if peekN(r, 4) == "do()" {
			r.Discard(4)
			return Token{Type: TypeDo, Value: "do()"}
		}
	}

	if b == 'm' {
		if peekN(r, 4) == "mul(" {
			result := Token{Type: TypeMultiply, Value: "mul"}
			r.Discard(4)

			value1 := next(r)
			if value1.Type != TypeInteger {
				return value1
			}
			comma := peek(r)
			if comma != ',' {
				return Token{Type: TypeUnknown, Value: string(comma)}
			}
			r.Discard(1)

			value2 := next(r)
			if value2.Type != TypeInteger {
				return value2
			}

			a, err := strconv.Atoi(value1.Value)
			if err != nil {
				ui.Die(err)
			}

			b, err := strconv.Atoi(value2.Value)
			if err != nil {
				ui.Die(err)
			}

			result.Value = fmt.Sprintf("%d", a*b)

			if peek(r) == ')' {
				r.Discard(1)
				return result
			}
		}
	}

	r.Discard(1)

	return Token{Type: TypeUnknown, Value: string(b)}
}

func (d day03) Part1(reader io.Reader) int {
	r := bufio.NewReader(reader)

	total := 0

outer:
	for {
		next := next(r)

		switch next.Type {
		case TypeEOF:
			break outer
		case TypeMultiply:
			v, err := strconv.Atoi(next.Value)
			if err != nil {
				ui.Die(err)
			}
			total += v
		default:
			// Do nothing
		}
	}

	return total
}

func (d day03) Part2(reader io.Reader) int {
	r := bufio.NewReader(reader)

	do := true
	total := 0

outer:
	for {
		next := next(r)

		switch next.Type {
		case TypeEOF:
			break outer
		case TypeDo:
			do = true
		case TypeDont:
			do = false
		case TypeMultiply:
			if !do {
				break
			}
			v, err := strconv.Atoi(next.Value)
			if err != nil {
				ui.Die(err)
			}
			total += v
		case TypeUnknown:
		case TypeInteger:
			// Do nothing
		default:
			ui.Die(fmt.Errorf("unexpected token type %s", next.Type))
		}
	}

	return total
}
