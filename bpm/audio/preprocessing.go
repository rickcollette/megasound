package audio

import (
	"math"
	"github.com/rickcollette/megasound/bpm/utils"
)

// LowPassFilter applies a simple low-pass filter to remove noise from samples.
func LowPassFilter(samples []float32, cutoff float64, rate int) []float32 {
	filtered := make([]float32, len(samples))
	alpha := cutoff / (cutoff + float64(rate))
	for i := 1; i < len(samples); i++ {
		filtered[i] = filtered[i-1] + float32(alpha)*(samples[i]-filtered[i-1])
	}
	return filtered
}

// ReadFloatArray processes PCM samples into an energy array.
func ReadFloatArray(samples []float32) []float32 {
	var v, n float64
	nrg := make([]float32, 0)

	for i := 0; i < len(samples); i++ {
		z := math.Abs(float64(samples[i]))
		if z > v {
			v += (z - v) / 8
		} else {
			v -= (v - z) / 512
		}

		n++
		if n == float64(utils.INTERVAL) { // Cast utils.INTERVAL to float64
			n = 0
			nrg = append(nrg, float32(v))
		}
	}

	return nrg
}