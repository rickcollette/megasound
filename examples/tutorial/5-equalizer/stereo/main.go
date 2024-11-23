package main

import (
	"math/rand"
	"time"

	"github.com/rickcollette/megasound"
	"github.com/rickcollette/megasound/effects"
	"github.com/rickcollette/megasound/speaker"
)

func noise() megasound.Streamer {
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

	eq := effects.NewEqualizer(noise(), sr, effects.StereoEqualizerSections{
		{
			Left:  effects.MonoEqualizerSection{F0: 200, Bf: 5, GB: 3, G0: 0, G: 8},
			Right: effects.MonoEqualizerSection{F0: 200, Bf: 5, GB: 3, G0: 0, G: -8},
		},
		{
			Left:  effects.MonoEqualizerSection{F0: 10000, Bf: 1000, GB: 3, G0: 0, G: 90},
			Right: effects.MonoEqualizerSection{F0: 10000, Bf: 1000, GB: 3, G0: 0, G: -90},
		},
	})

	speaker.Play(eq)
	select {}
}
