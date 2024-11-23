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
	f, err := os.Open("gunshot.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/60))

	buffer := megasound.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	for {
		fmt.Print("Press [ENTER] to fire a gunshot! ")
		fmt.Scanln()

		shot := buffer.Streamer(0, buffer.Len())
		speaker.Play(shot)
	}
}