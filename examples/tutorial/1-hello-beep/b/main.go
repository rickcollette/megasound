package main

import (
	"log"
	"os"
	"time"

	"github.com/rickcollette/megasound"
	"github.com/rickcollette/megasound/mp3"
	"github.com/rickcollette/megasound/speaker"
)

func main() {
	f, err := os.Open("../Lame_Drivers_-_01_-_Frozen_Egg.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	sr := format.SampleRate * 2
	speaker.Init(sr, sr.N(time.Second/10))

	resampled := megasound.Resample(4, format.SampleRate, sr, streamer)

	done := make(chan bool)
	speaker.Play(megasound.Seq(resampled, megasound.Callback(func() {
		done <- true
	})))

	<-done
}
