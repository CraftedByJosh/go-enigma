package types

import "strings"

type Position int

func (p Position) ToString() string {
	return string((int(p) + 97))
}

const (
	A Position = 0
	B Position = 1
	C Position = 2
	D Position = 3
	E Position = 4
	F Position = 5
	G Position = 6
	H Position = 7
	I Position = 8
	J Position = 9
	K Position = 10
	L Position = 11
	M Position = 12
	N Position = 13
	O Position = 14
	P Position = 15
	Q Position = 16
	R Position = 17
	S Position = 18
	T Position = 19
	U Position = 20
	V Position = 21
	W Position = 22
	X Position = 23
	Y Position = 24
	Z Position = 25
)

// wrap normalizes a given Position (or int) to stay within 0–25
// It ensures correct wrapping behavior for both positive and negative input values.
//
// Examples:
//   wrap(27)  => 1   (27 mod 26 = 1)
//   wrap(-1)  => 25  (-1 mod 26 = 25, i.e. Z)
//   wrap(0)   => 0   (A)
//   wrap(26)  => 0   (wraps to A again)
func WrapPosition(p Position) Position {
    return (p%26 + 26) % 26
}

func StringToPositionSlice(s string) []Position {
	positions := make([]Position, len(s))
	upperString := strings.ToUpper(s)

	for i, char := range upperString {
		positions[i] = Position(char - 'A')
	}

	return positions
}
