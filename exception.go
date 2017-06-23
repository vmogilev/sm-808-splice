package main

import "fmt"

// Exception is a helper type to standartize on errors
type exception string

var excepErrors = map[exception]string{
	tempoRange:       "Tempo range should be between 60 and 128 BPM",
	tempoNotNumber:   "Tempo should be a number",
	songTitleTooLong: fmt.Sprintf("Song Title should be less than %d", maxSongTitle),
}

const tempoRange = exception("tempoRange")
const tempoNotNumber = exception("tempoNotNumber")
const songTitleTooLong = exception("songTitleTooLong")

func (e exception) Error() string {
	return fmt.Sprintf("%s", excepErrors[e])
}
