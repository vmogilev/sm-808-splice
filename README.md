# Splice sm-808 Implementation

Here's my implementation of the Splice SM-808 [practice exercise](https://github.com/splicers/sm-808)

![rotating pattern grid with a moving cursor underneath](/sm-808-animation.gif?raw=true)

Few notes about this implementation:

* It's a rotating pattern grid with a moving cursor underneath it to indicate which step the drum machine is currently on.  It's using terminal escape sequences to position the cursor - see `song.go` - `playStep()`.
* Supports mixing patterns of different durations (8, 16, 32 steps) and repeats the shorter pattern matching the length of the longest duration.
* Supports specifying a velocity via `Pattern.Beats` which is a map of steps to velocity.  See example below - where a `snare` has a velocity of `3` at `13th` step, while all other patterns and steps have a default velocity of `1`:

        song := NewSong(title, bpm)
        kick := map[int]int{1: 1, 5: 1}
        snare := map[int]int{5: 1, 13: 3}
        hihat := map[int]int{3: 1, 7: 1}

        song.AddPattern("Kick", kick, "Kick.aif")
        song.AddPattern("Snare", snare, "Snare.aif")
        song.AddPattern("HiHat", hihat, "HiHat.aif")

* I ran out of time implementing outputting the sound.  See `play.no_go` for my attempt, which is just a copy from [gordonklaus/portaudio](https://github.com/gordonklaus/portaudio) - I got it to play the patterns but they are not in synch with the beat (there is too much lag).

End.
