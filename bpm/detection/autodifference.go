package detection

import (
	"github.com/rickcollette/megasound/bpm/utils"
	"math"
	"sync"
)

// sample extracts a value from the energy array at a given offset.
func sample(nrg []float32, offset float64) float64 {
	i := int64(math.Floor(offset))
	if i >= 0 && i < int64(len(nrg)) {
		return float64(nrg[i])
	}
	return 0.0
}

// autodifference computes the energy difference for a given interval.
func autodifference(nrg []float32, interval float64) float64 {
	beats := []float64{-32, -16, -8, -4, -2, -1, 1, 2, 4, 8, 16, 32}
	result := make(chan float64, len(beats))
	var wg sync.WaitGroup

	for _, beat := range beats {
		wg.Add(1)
		go func(beat float64) {
			defer wg.Done()
			y := sample(nrg, float64(len(nrg)/2)+beat*interval)
			diff := math.Abs(y) / math.Abs(beat)
			result <- diff
		}(beat)
	}

	wg.Wait()
	close(result)

	// Aggregate results
	var totalDiff float64
	for diff := range result {
		totalDiff += diff
	}
	return totalDiff
}

// bpmToInterval converts BPM to energy sampling interval.
func bpmToInterval(bpm float64) float64 {
	return float64(utils.RATE) / (bpm / 60) / float64(utils.INTERVAL)
}

// intervalToBpm converts an energy sampling interval back to BPM.
func intervalToBpm(interval float64) float64 {
	return (float64(utils.RATE) / (interval * float64(utils.INTERVAL))) * 60
}
