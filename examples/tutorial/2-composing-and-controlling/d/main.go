package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rickcollette/megasound"
	"github.com/rickcollette/megasound/effects"
	"github.com/rickcollette/megasound/mp3"
	"github.com/rickcollette/megasound/speaker"
)

func main() {
	f, err := os.Open("../Miami_Slice_-_04_-_Step_Into_Me.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ctrl := &megasound.Ctrl{Streamer: megasound.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speedy := megasound.ResampleRatio(4, 1, volume)
	speaker.Play(speedy)

	for {
		fmt.Print("Press [ENTER] to pause/resume. ")
		fmt.Scanln()

		speaker.Lock()
		ctrl.Paused = !ctrl.Paused
		volume.Volume += 0.5
		speedy.SetRatio(speedy.Ratio() + 0.1)
		speaker.Unlock()
	}
}