package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Pattern - a pattern can be added to a song, each pattern can have
// any number of beats assigned as a map of step ID to int values measuring
// velocity as follows:
//
//      step[1] = 0 - no beat - skip
//      step[2] = 1 - one beat at velocity 1
//      step[3] = 2 - one beat at velocity 2
//      step[n] = n - one beat at velocity n
//      ...
//
// File is an optinal sound file to play
type Pattern struct {
	Name  string
	Beats map[int]int // Map of beats velocity for each step of a pattern
	File  string      // Optional sound file to play
}

// Song is a collection of patterns under one name and tempo (measured in bpm)
type Song struct {
	Name     string
	Patterns []Pattern
	Tempo    int
}

const defaultBpm = 128
const defaultSongName = "Four-on-the-floor"

func main() {
	var bpm int
	var err error
	for {
		fmt.Printf("\tEnter Tempo (bpm):")
		if bpm, err = parseTempo(os.Stdin); err == nil {
			break
		}
		fmt.Printf("\t\t ** %s **\n", err.Error())
	}
	fmt.Printf("Using %d BPM ...\n", bpm)
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
	pat := Pattern{
		Name:  name,
		Beats: beats,
	}
	s.Patterns = append(s.Patterns, pat)
}

func parseTempo(r io.Reader) (int, error) {
	var bpm int
	var err error
	var bpmStr string

	reader := bufio.NewReader(r)
	bpmStr, _ = reader.ReadString('\n')
	bpmStr = strings.TrimSpace(bpmStr)

	if len(bpmStr) == 0 {
		return defaultBpm, nil
	}

	bpm, err = strconv.Atoi(bpmStr)
	if err != nil {
		return 0, tempoNotNumber
	}

	if bpm < 60 || bpm > 128 {
		return 0, tempoRange
	}

	return bpm, nil
}
