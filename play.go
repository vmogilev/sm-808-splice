package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/gordonklaus/portaudio"
)

func play(fileName string) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	f, err := os.Open(fileName)
	chk(err)
	defer f.Close()

	id, data, err := readChunk(f)
	chk(err)
	if id.String() != "FORM" {
		fmt.Printf("bad file format - 1: *%s*\n", id.String())
		return
	}
	_, err = data.Read(id[:])
	chk(err)
	if id.String() != "AIFF" {
		fmt.Println("bad file format - 2")
		return
	}
	var c commonChunk
	var audio io.Reader
	for {
		id, chunk, err := readChunk(data)
		if err == io.EOF {
			break
		}
		chk(err)
		switch id.String() {
		case "COMM":
			chk(binary.Read(chunk, binary.BigEndian, &c))
		case "SSND":
			chunk.Seek(8, 1) //ignore offset and block
			audio = chunk
			// default:
			// 	fmt.Printf("ignoring unknown chunk '%s'\n", id)
		}
	}

	//assume 44100 sample rate, mono, 32 bit

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int32, 8192)
	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(out), &out)
	fmt.Printf("%s%s", moveCursor(1, 60), fileName)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	defer stream.Stop()
	i := 0
	for remaining := int(c.NumSamples); remaining > 0; remaining -= len(out) {
		fmt.Printf("%s%d", moveCursor(2, 60), remaining)
		i++
		if len(out) > remaining {
			out = out[:remaining]
		}
		err := binary.Read(audio, binary.BigEndian, out)
		if err == io.EOF {
			break
		}
		chk(err)
		chk(stream.Write())
		select {
		case <-sig:
			return
		default:
		}
	}
}

func readChunk(r readerAtSeeker) (id ID, data *io.SectionReader, err error) {
	_, err = r.Read(id[:])
	if err != nil {
		return
	}
	var n int32
	err = binary.Read(r, binary.BigEndian, &n)
	if err != nil {
		return
	}
	off, _ := r.Seek(0, 1)
	data = io.NewSectionReader(r, off, int64(n))
	_, err = r.Seek(int64(n), 1)
	return
}

type readerAtSeeker interface {
	io.Reader
	io.ReaderAt
	io.Seeker
}

type ID [4]byte

func (id ID) String() string {
	return string(id[:])
}

type commonChunk struct {
	NumChans      int16
	NumSamples    int32
	BitsPerSample int16
	SampleRate    [10]byte
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
