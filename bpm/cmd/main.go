package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rickcollette/megasound/bpm/audio"
	"github.com/rickcollette/megasound/bpm/detection"
	"github.com/rickcollette/megasound/bpm/utils"
)

var (
	min                 = flag.Float64("min", 120, "min BPM you are expecting")
	max                 = flag.Float64("max", 200, "max BPM you are expecting")
	progressive         = flag.Bool("progressive", false, "Print the BPM for every period")
	progressiveInterval = flag.Int("interval", 10, "How many seconds for every progressive chunk printed")
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	filePath := flag.Arg(0)
	log.Printf("Processing file: %s\n", filePath)

	// Get metadata from the audio file
	metadata, err := audio.GetMetadata(filePath)
	if err != nil {
		log.Fatalf("Failed to extract audio metadata: %v", err)
	}
	log.Printf("Sample Rate: %d Hz, Channels: %d\n", metadata.Rate, metadata.Channels)

	// Set dynamic RATE and INTERVAL
	utils.RATE = metadata.Rate
	utils.INTERVAL = audio.CalculateInterval(metadata.Rate, 128)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer file.Close()

	in := make(chan float32)
	out := make(chan float32)

	log.Println("Starting progressive energy level processing.")
	go audio.ProgressivelyReadFloatArray(in, out)

	done := make(chan bool)
	go readProgressiveVars(out, done, *progressive, *progressiveInterval)

	// Read data from file and send to the input channel
	for {
		var f float32
		err = binary.Read(file, binary.LittleEndian, &f)
		if err != nil {
			break
		}
		in <- f
	}
	close(in)

	// Wait for the processing to complete
	<-done
	log.Println("Processing complete.")
}


func readProgressiveVars(input chan float32, done chan bool, progressive bool, pint int) {
	if progressive {
		chunkSize := detection.CalcChunkLen(pint)
		nrg := make([]float32, 0, chunkSize)
		for f := range input {
			nrg = append(nrg, f)
			if len(nrg) == chunkSize {
				bpm := detection.ScanForBpm(nrg, *min, *max, 1024, 1024)
				fmt.Printf("Progressive BPM: %.2f\n", bpm)
				nrg = make([]float32, 0, chunkSize)
			}
		}
		done <- true
	} else {
		nrg := make([]float32, 0)
		for f := range input {
			nrg = append(nrg, f)
		}
		bpm := detection.ScanForBpm(nrg, *min, *max, 1024, 1024)
		fmt.Printf("Final BPM: %.2f\n", bpm)
		done <- true
	}
}
