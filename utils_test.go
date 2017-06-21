package main

import (
	"testing"
)

var testMoveCursorCases = []struct {
	row    int
	column int
	want   string
}{
	{1, 1, "\033[1;1H"},
	{11, 100, "\033[11;100H"},
}

var testHeaderCases = []struct {
	column int
	step   int
	want   string
}{
	{11, 1, "\033[1;11H 1 "},
	{41, 64, "\033[1;41H 8 "},
	{21, 38, "\033[1;21H 6 "},
}

var testFooterCases = []struct {
	maxColumns int
	column     int
	lastRow    int
	want       string
}{
	{42, 35, 6, "\033[6;1H                                  *       "},
}

func TestFooter(t *testing.T) {
	for _, test := range testFooterCases {
		if got := footer(test.maxColumns, test.column, test.lastRow); got != test.want {
			t.Fatalf("footer(maxColumns=%d, column=%d, lastRow=%d) Want: %q Got: %q\n", test.maxColumns, test.column, test.lastRow, test.want, got)
		}
	}
}

func TestHeader(t *testing.T) {
	for _, test := range testHeaderCases {
		if got := header(test.column, test.step); got != test.want {
			t.Fatalf("header(column=%d, step=%d) Want: %q Got: %q\n", test.column, test.step, test.want, got)
		}
	}
}

func TestMoveCursorCases(t *testing.T) {
	for _, test := range testMoveCursorCases {
		if got := moveCursor(test.row, test.column); got != test.want {
			t.Fatalf("moveCursor(row=%d, column=%d) Want: %q Got: %q\n", test.row, test.column, test.want, got)
		}
	}
}
