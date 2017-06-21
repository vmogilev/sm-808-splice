package main

import (
	"fmt"
	"sort"
	"time"
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

// MaxPatDur - gets the max length of a pattern
func (s *Song) MaxPatDur() int {
	pats := make([]int, len(s.Patterns))
	i := 0
	for _, p := range s.Patterns {
		pats[i] = p.Duration
		i++
	}
	sort.Ints(pats)
	return pats[len(pats)-1]
}

// AddPattern - adds a specific beat pattern to existing song
func (s *Song) AddPattern(name string, beats map[int]int, file string) {
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
		File:     file,
	}
	s.Patterns = append(s.Patterns, pat)
}

// Play - prints/plays all patterns at specific step of a song
func (s *Song) Play(step int) (out string, column int, dur float64) {
	if step == 1 {
		out = s.printHeaders()
	}

	beats, column, dur := s.playStep(step)
	return out + beats, column, dur
}

func (s *Song) printHeaders() (out string) {
	out = fmt.Sprintf("%s%7s   \n", moveCursor(1, 1), " ")
	for i, pat := range s.Patterns {
		out += fmt.Sprintf("%s%7s: |\n", moveCursor(i+2, 1), pat.Name)
	}
	return out
}

func (s *Song) playStep(step int) (out string, column int, dur float64) {
	var headerLength = 10

	// normalize step over maximum pattern length
	maxDur := s.MaxPatDur()
	stepNorm := step % maxDur
	if stepNorm == 0 {
		stepNorm = maxDur
	}

	// The escape seq: \033[1;1H is produced by moveCursor(row,column)
	//                      ^ ^
	//                      ^ col
	//                      ^
	//                      row
	//
	// in the above example it's job is to move cursor to row=1, col=1
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
	case stepNorm == 1:
		column = headerLength + 1
	case stepNorm >= 2:
		column = headerLength + (2 * stepNorm) - 1
	}

	dur = (((60.0 / float64(s.Tempo)) * 4.0) / 8.0)
	microSecs := time.Duration(dur*1000000) * time.Microsecond

	var lastRow int
	var xOrUnderscore string
	for row, pat := range s.Patterns {
		key := stepNorm % pat.Duration
		if key == 0 {
			key = pat.Duration
		}
		if _, ok := pat.Beats[key]; ok {
			xOrUnderscore = "X"
			// if pat.File != "" {
			// 	play("beats" + string(filepath.Separator) + pat.File)
			// }
		} else {
			xOrUnderscore = "_"
		}
		out += fmt.Sprintf("%s%s|", moveCursor(row+2, column), xOrUnderscore)
		lastRow = row + 1
		//fmt.Printf("\033[%d;%dH%d\n", row+10, 1, step%pat.Duration)
	}

	// I also want a cursor (*) under the current beat column
	// because we wrap over the steps when they get > maxDur
	// and a heading for each step (normalized in 8 steps)
	//
	// Example:
	//
	//         1 2 3 4 5 6 7 8 1 2 3 4 5 6 7 8
	//  Kick: |X|_|_|_|X|_|_|_|X|_|_|_|X|_|_|_|
	// Snare: |_|_|_|_|X|_|_|_|_|_|_|_|X|_|_|_|
	// HiHat: |_|_|X|_|_|_|X|_|_|_|X|_|_|_|X|_|
	// HiTom: |_|_|_|_|_|X|_|_|_|_|_|X|_|_|_|X|
	// 			     *
	//
	// >> Step: 44
	// >> Column: 33
	//
	maxColumns := headerLength + (2 * maxDur)
	head := header(column-1, stepNorm)
	foot := footer(maxColumns, column, lastRow+2)
	time.Sleep(microSecs)
	return head + out + foot, column, dur
}
