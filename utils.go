package main

import (
	"fmt"
	"strings"
)

// moveCursor - produces escape seq: \033[1;1H
//                                        ^ ^
//                                        ^ col
//                                        ^
//                                        row
// in the above example it's job is to move cursor to row=1, col=1
// see: https://en.wikipedia.org/wiki/ANSI_escape_code#Sequence_elements
//      https://stackoverflow.com/questions/15442292/golang-how-to-have-an-inplace-string-that-updates-at-stdout
func moveCursor(row int, column int) string {
	return fmt.Sprintf("\033[%d;%dH", row, column)
}

func header(column int, step int) string {
	stepIn8 := step % 8
	if stepIn8 == 0 {
		stepIn8 = 8
	}
	return fmt.Sprintf("%s %d ", moveCursor(1, column), stepIn8)
}

func footer(maxColumns int, column int, lastRow int) string {
	leftPad := strings.Repeat(" ", column-1)
	rightPad := strings.Repeat(" ", maxColumns-column)
	return fmt.Sprintf("%s%s", moveCursor(lastRow, 1), leftPad+"*"+rightPad)
}
