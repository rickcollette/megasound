package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rickcollette/megasound"
	"github.com/rickcollette/megasound/generators"
	"github.com/rickcollette/megasound/speaker"
)

func usage() {
	fmt.Printf("usage: %s freq\n", os.Args[0])
	fmt.Println("where freq must be a float between 1 and 24000")
	fmt.Println("24000 because samplerate of 48000 is hardcoded")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	f, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		usage()
		return
	}

	sr := megasound.SampleRate(48000)
	speaker.Init(sr, 4800)

	sine, err := generators.SineTone(sr, f)
	if err != nil {
		panic(err)
	}

	triangle, err := generators.TriangleTone(sr, f)
	if err != nil {
		panic(err)
	}

	square, err := generators.SquareTone(sr, f)
	if err != nil {
		panic(err)
	}

	sawtooth, err := generators.SawtoothTone(sr, f)
	if err != nil {
		panic(err)
	}

	sawtoothReversed, err := generators.SawtoothToneReversed(sr, f)
	if err != nil {
		panic(err)
	}

	// Play 2 seconds of each tone
	two := sr.N(2 * time.Second)

	ch := make(chan struct{})
	sounds := []megasound.Streamer{
		megasound.Callback(print("sine")),
		megasound.Take(two, sine),
		megasound.Callback(print("triangle")),
		megasound.Take(two, triangle),
		megasound.Callback(print("square")),
		megasound.Take(two, square),
		megasound.Callback(print("sawtooth")),
		megasound.Take(two, sawtooth),
		megasound.Callback(print("sawtooth reversed")),
		megasound.Take(two, sawtoothReversed),
		megasound.Callback(func() {
			ch <- struct{}{}
		}),
	}
	speaker.Play(megasound.Seq(sounds...))

	<-ch
}

// a simple clousure to wrap fmt.Println
func print(s string) func() {
	return func() {
		fmt.Println(s)
	}
}
