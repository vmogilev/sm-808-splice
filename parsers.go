package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const defaultBpm = 128
const defaultSongName = "Four-on-the-floor"
const maxSongTitle = 30
const maxBpm = 128
const minBpm = 30

func parseTitle(r io.Reader) (string, error) {
	var title string

	reader := bufio.NewReader(r)
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if len(title) > maxSongTitle {
		return "", songTitleTooLong
	}

	if title == "" {
		title = defaultSongName
	}

	return title, nil
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

	if bpm < minBpm || bpm > maxBpm {
		return 0, tempoRange
	}

	return bpm, nil
}
