package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	var title string
	var bpm int
	var err error

	for {
		fmt.Printf("\tEnter Song Title (Default: %s): ", defaultSongName)
		if title, err = parseTitle(os.Stdin); err == nil {
			break
		}
		fmt.Printf("\t\t ** %s **\n", err.Error())
	}
	fmt.Printf("\t> Your cool song is \"%s\"\n", title)

	for {
		fmt.Printf("\tEnter Tempo (Default: %dbpm): ", defaultBpm)
		if bpm, err = parseTempo(os.Stdin); err == nil {
			break
		}
		fmt.Printf("\t\t ** %s **\n", err.Error())
	}
	fmt.Printf("\t> Using %d BPM ...\n", bpm)

	song := NewSong(title, bpm)
	kick := map[int]int{1: 1, 5: 1}
	snare := map[int]int{5: 1, 13: 1}
	hihat := map[int]int{3: 1, 7: 1}
	//hitom := map[int]int{6: 1, 12: 1, 16: 1}
	song.AddPattern("Kick", kick, "Kick.aif")
	song.AddPattern("Snare", snare, "Snare.aif")
	song.AddPattern("HiHat", hihat, "HiHat.aif")
	//song.AddPattern("HiTom", hitom)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 1; i <= song.MaxPatDur()*2; i++ {
		beats, column := song.Play(i)
		fmt.Printf("%s", beats)
		fmt.Printf("\n\n>> Step: %d\n", i)
		fmt.Printf(">> Column: %d\n", column)
		sec := (((60.0 / float64(bpm)) * 4.0) / 8.0)
		fmt.Printf(">> Duration: %.2f sec\n", sec)
		dur := time.Duration(sec*1000000) * time.Microsecond
		time.Sleep(dur)
	}

}
