package main

import (
	"bytes"
	"testing"
)

var testSongTitleCases = []struct {
	in  string
	out string
	ok  error
}{
	{"\n", defaultSongName, nil},
	{"Four On The Floor\n", "Four On The Floor", nil},
	{"Four On\n The Floor\n", "Four On", nil},
	{"123456789012345678901234567890123\n", "", songTitleTooLong},
}

var testTempoCases = []struct {
	in  string
	out int
	ok  error
}{
	{
		"1\n",
		0,
		tempoRange,
	},
	{
		"300\n",
		0,
		tempoRange,
	},
	{
		"10\n",
		0,
		tempoRange,
	},
	{
		"60\n",
		60,
		nil,
	},
	{
		"\n",
		128,
		nil,
	},
	{
		"ABC\n",
		0,
		tempoNotNumber,
	},
}

func TestTempo(t *testing.T) {
	for _, test := range testTempoCases {
		if out, ok := parseTempo(bytes.NewBufferString(test.in)); out != test.out || ok != test.ok {
			t.Fatalf("In: %s Want: %d | %s\nGot: %d | %s\n", test.in, test.out, test.ok, out, ok)
		}
	}
}

func TestSongTitle(t *testing.T) {
	for _, test := range testSongTitleCases {
		if out, ok := parseTitle(bytes.NewBufferString(test.in)); out != test.out || ok != test.ok {
			t.Fatalf("In: %s Want: %s | %s\nGot: %s | %s\n", test.in, test.out, test.ok, out, ok)
		}
	}
}
