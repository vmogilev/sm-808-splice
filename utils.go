package main

import "fmt"

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
