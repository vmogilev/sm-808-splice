package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Pattern - a pattern can be added to a song, each pattern has 8 beat slots
// that can be filled with int values used as follows:
//
//      0 - no beat - skip
//      1 - output one bit at velocity 1
//      2 - output one bit at velocity 2
//      n - output one bit at velocity n ...
//
// File is an optinal sound file to play
type Pattern struct {
	Name  string
	Beats []int  // beats in 8 step measure
	File  string // Optional sound file to play
}

// Song is a collection of patterns under one name and tempo (measured in bpm)
type Song struct {
	Name     string
	Patterns []Pattern
	Tempo    int
}

func main() {
	var bpm int
	var err error
	for {
		if bpm, err = parseTempo(os.Stdin); err == nil {
			break
		}
		fmt.Printf("\t\t ** %s **\n", err.Error())
	}
	fmt.Println(bpm)
}

func parseTempo(r io.Reader) (int, error) {
	var bpm = 128
	var err error
	var bpmStr string

	reader := bufio.NewReader(r)
	fmt.Printf("\tEnter Tempo (bpm):")
	bpmStr, _ = reader.ReadString('\n')
	bpmStr = strings.TrimSpace(bpmStr)

	if len(bpmStr) == 0 {
		fmt.Printf("\t\tusing default at %d BPM ...", bpm)
		return bpm, nil
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
