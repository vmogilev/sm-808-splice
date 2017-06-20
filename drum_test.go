package main

import (
	"bytes"
	"reflect"
	"testing"
)

var testSongTitleCases = []struct {
	in  string
	out string
	ok  error
}{
	{"\n", defaultSongName, nil},
	{"Four On The Floor\n", "Four On The Floor", nil},
	{"Four On\n The Floor\n", "Four On", nil},
	{"123456789012345678901234567890123\n", "", songTitleTooLong},
}

var testTempoCases = []struct {
	in  string
	out int
	ok  error
}{
	{
		"1\n",
		0,
		tempoRange,
	},
	{
		"300\n",
		0,
		tempoRange,
	},
	{
		"10\n",
		0,
		tempoRange,
	},
	{
		"60\n",
		60,
		nil,
	},
	{
		"\n",
		128,
		nil,
	},
	{
		"ABC\n",
		0,
		tempoNotNumber,
	},
}

var testNewSongCases = []struct {
	title string
	tempo int
	song  *Song
}{
	{
		"",
		60,
		&Song{Name: defaultSongName, Tempo: 60},
	},
	{
		"November Rain",
		128,
		&Song{Name: "November Rain", Tempo: 128},
	},
}

var testPatternCases = []struct {
	songIn  *Song
	name    string
	beats   map[int]int
	songOut *Song
}{
	{
		&Song{Name: "One", Tempo: 70},
		"HiHat",
		map[int]int{15: 1, 7: 1, 11: 1, 3: 1},
		&Song{
			Name:  "One",
			Tempo: 70,
			Patterns: []Pattern{
				Pattern{Name: "HiHat", Beats: map[int]int{15: 1, 7: 1, 11: 1, 3: 1}, Duration: 16},
			},
		},
	},
	{
		&Song{Name: "Two", Tempo: 70},
		"HiHat",
		map[int]int{16: 1, 7: 1, 11: 1, 3: 1},
		&Song{
			Name:  "Two",
			Tempo: 70,
			Patterns: []Pattern{
				Pattern{Name: "HiHat", Beats: map[int]int{16: 1, 7: 1, 11: 1, 3: 1}, Duration: 16},
			},
		},
	},
	{
		&Song{
			Name:  "With One",
			Tempo: 70,
			Patterns: []Pattern{
				Pattern{Name: "HiHat", Beats: map[int]int{3: 1, 7: 1, 11: 1, 15: 1}, Duration: 16},
			},
		},
		"Snare Drum",
		map[int]int{5: 1, 13: 1},
		&Song{
			Name:  "With One",
			Tempo: 70,
			Patterns: []Pattern{
				Pattern{Name: "HiHat", Beats: map[int]int{3: 1, 7: 1, 11: 1, 15: 1}, Duration: 16},
				Pattern{Name: "Snare Drum", Beats: map[int]int{5: 1, 13: 1}, Duration: 16},
			},
		},
	},
}

func TestSongTitle(t *testing.T) {
	for _, test := range testSongTitleCases {
		if out, ok := parseTitle(bytes.NewBufferString(test.in)); out != test.out || ok != test.ok {
			t.Fatalf("In: %s Want: %s | %s\nGot: %s | %s\n", test.in, test.out, test.ok, out, ok)
		}
	}
}

func TestPatterns(t *testing.T) {
	for _, test := range testPatternCases {
		test.songIn.AddPattern(test.name, test.beats)
		if !reflect.DeepEqual(test.songIn, test.songOut) {
			t.Fatalf("In: %v\nGot: %v\nWant: %v", test, test.songIn, test.songOut)
		}
	}
}

func TestTempo(t *testing.T) {
	for _, test := range testTempoCases {
		if out, ok := parseTempo(bytes.NewBufferString(test.in)); out != test.out || ok != test.ok {
			t.Fatalf("In: %s Want: %d | %s\nGot: %d | %s\n", test.in, test.out, test.ok, out, ok)
		}
	}
}

func TestSong(t *testing.T) {
	for _, test := range testNewSongCases {
		if song := NewSong(test.title, test.tempo); !reflect.DeepEqual(song, test.song) {
			t.Fatalf("Title: %s | Tempo: %d\nGot: %v | Want: %v", test.title, test.tempo, song, test.song)
		}
	}
}
