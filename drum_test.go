package main

import (
	"bytes"
	"testing"
)

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
