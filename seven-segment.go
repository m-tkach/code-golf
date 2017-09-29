package main

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var segments = [][]string{
	{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "},
	{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"},
	{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"},
}

func sevenSegment() (input, output string) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	digits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	// Shuffle the digits
	for i := range digits {
		j := random.Intn(i + 1)
		digits[i], digits[j] = digits[j], digits[i]
	}

	input = string(digits)

	for row := 0; row < 3; row++ {
		for _, digit := range digits {
			output += segments[row][digit-'0']
		}

		output += "\n"
	}

	output = strings.TrimRightFunc(output, unicode.IsSpace)

	return
}
