package main

import (
	"fmt"
	"sort"
)

// Pattern - a pattern can be added to a song, each pattern can have
// any number of beats assigned as a map of step ID to int values measuring
// velocity as follows:
//
//      step[2] = 1 - one beat at velocity 1
//      step[3] = 2 - one beat at velocity 2
//      step[n] = n - one beat at velocity n
//      ...
//
// File is an optinal sound file to play
type Pattern struct {
	Name     string
	Beats    map[int]int // Map of beats velocity for each step of a pattern
	Duration int
	File     string // Optional sound file to play
}

// Song is a collection of patterns under one name and tempo (measured in bpm)
type Song struct {
	Name     string
	Tempo    int
	Patterns []Pattern
}

// NewSong - creates a new song with a default title if none provided
func NewSong(t string, tempo int) *Song {
	title := t
	if title == "" {
		title = defaultSongName
	}
	return &Song{
		Name:  title,
		Tempo: tempo,
	}
}

// AddPattern - adds a specific beat pattern to existing song
func (s *Song) AddPattern(name string, beats map[int]int) {
	keys := make([]int, len(beats))
	i := 0
	for k := range beats {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	maxKey := keys[len(keys)-1]

	var duration int
	for _, d := range []int{32, 16, 8} {
		if d/2 > d-maxKey {
			duration = d
			break
		}
	}

	pat := Pattern{
		Name:     name,
		Beats:    beats,
		Duration: duration,
	}
	s.Patterns = append(s.Patterns, pat)
}

// Play - prints/plays all patterns at specific step of a song
func (s *Song) Play(step int) (out string, column int) {
	if step == 1 {
		out = s.printHeaders()
	}

	beats, column := s.playStep(step)
	return out + beats, column
}

func (s *Song) printHeaders() (out string) {
	for _, pat := range s.Patterns {
		out += fmt.Sprintf("%7s: |\n", pat.Name)
	}
	return out
}

func (s *Song) playStep(step int) (out string, column int) {
	var headerLength = 10
	format := "\033[%d;%dH%s\033[%d;%dH|"

	// The escape seq: \033[1;1H = moves cursor to row=1, col=1
	//                      ^ ^
	//                      ^ col
	//                      ^
	//                      row
	//
	// see: https://en.wikipedia.org/wiki/ANSI_escape_code#Sequence_elements
	//      https://stackoverflow.com/questions/15442292/golang-how-to-have-an-inplace-string-that-updates-at-stdout
	//
	// So in the above format we are escaping twice:
	//      1) the single char string (_ or X)
	//      2) and the seperator (|)
	//
	// Each column is caluculated as follows:
	//
	// 1234567890123456789012345678901  <-- columns
	//           1 2 3 4 5 6 7 8 9 0 1  <-- steps
	//    Kick: |X|_|_|_|X|_|_|_|X|_|_|
	//   Snare: |_|_|_|_|X|_|_|_|_|_|_|
	//   HiHat: |_|_|X|_|_|_|X|_|_|_|X|
	//                               ^
	// >> Step: 11
	// >> Column: 31
	//
	// 1) each step has two chars in it `_|` or `X|`
	// 2) so we have to multiply step*2, add header (10) and
	//          then step one back to get to the current column
	// 3) in this example we are at step 11 / column 31
	//          (11*2)+10-1=31
	//
	switch {
	case step == 1:
		column = headerLength + 1
	case step >= 2:
		column = headerLength + (2 * step) - 1
	}

	for row, pat := range s.Patterns {
		key := step % (pat.Duration)
		if key == 0 {
			key = pat.Duration
		}
		if _, ok := pat.Beats[key]; ok {
			out += fmt.Sprintf(format, row+1, column, "X", row+1, column+1)
		} else {
			out += fmt.Sprintf(format, row+1, column, "_", row+1, column+1)
		}
		//fmt.Printf("\033[%d;%dH%d\n", row+10, 1, step%pat.Duration)
	}
	return out, column
}
