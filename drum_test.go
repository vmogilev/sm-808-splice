package main

import (
	"bytes"
	"reflect"
	"testing"
)

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
