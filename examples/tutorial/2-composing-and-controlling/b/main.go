package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rickcollette/megasound"
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

	loop := megasound.Loop(3, streamer)
	fast := megasound.ResampleRatio(4, 5, loop)

	done := make(chan bool)
	speaker.Play(megasound.Seq(fast, megasound.Callback(func() {
		done <- true
	})))

	for {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			speaker.Lock()
			fmt.Println(format.SampleRate.D(streamer.Position()).Round(time.Second))
			speaker.Unlock()
		}
	}
}
