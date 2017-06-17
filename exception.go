package main

import "fmt"

// Exception is a helper type to standartize on errors
type exception string

var excepErrors = map[exception]string{
	"tempoRange":     "Tempo range should be between 60 and 128 BPM",
	"tempoNotNumber": "Tempo should be a number",
}

const tempoRange = exception("tempoRange")
const tempoNotNumber = exception("tempoNotNumber")

func (e exception) Error() string {
	return fmt.Sprintf("%s", excepErrors[e])
}
