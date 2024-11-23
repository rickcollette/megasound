package main

import (
	"math/rand"
	"time"

	"github.com/rickcollette/megasound"
	"github.com/rickcollette/megasound/speaker"
)

func Noise() megasound.Streamer {
	return megasound.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			samples[i][0] = rand.Float64()*2 - 1
			samples[i][1] = rand.Float64()*2 - 1
		}
		return len(samples), true
	})
}

func main() {
	sr := megasound.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(Noise())
	select {}
}
