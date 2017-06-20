package main

import (
	"reflect"
	"testing"
)

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

var testSongMaxDurCases = []struct {
	song *Song
	want int
}{
	{
		&Song{
			Name:  "One",
			Tempo: 70,
			Patterns: []Pattern{
				Pattern{Name: "1", Beats: map[int]int{15: 1, 7: 1, 11: 1, 3: 1}, Duration: 16},
				Pattern{Name: "2", Beats: map[int]int{1: 1, 2: 1, 6: 1}, Duration: 8},
			},
		},
		16,
	},
	{
		&Song{
			Name:  "Two",
			Tempo: 70,
			Patterns: []Pattern{
				Pattern{Name: "1", Beats: map[int]int{15: 1, 7: 1, 11: 1, 3: 1}, Duration: 16},
				Pattern{Name: "2", Beats: map[int]int{1: 1, 2: 1, 6: 1}, Duration: 8},
				Pattern{Name: "3", Beats: map[int]int{1: 1, 2: 1, 6: 1, 28: 1}, Duration: 32},
			},
		},
		32,
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

func TestSongMaxDur(t *testing.T) {
	for _, test := range testSongMaxDurCases {
		if out := test.song.MaxPatDur(); out != test.want {
			t.Fatalf("In: %v Want: %d Got: %d", test.song, test.want, out)
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

func TestPatterns(t *testing.T) {
	for _, test := range testPatternCases {
		test.songIn.AddPattern(test.name, test.beats)
		if !reflect.DeepEqual(test.songIn, test.songOut) {
			t.Fatalf("In: %v\nGot: %v\nWant: %v", test, test.songIn, test.songOut)
		}
	}
}
