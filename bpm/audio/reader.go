package audio

import (
	"math"
	"github.com/rickcollette/megasound/bpm/utils"
)

// ProgressivelyReadFloatArray processes PCM data into energy levels for BPM analysis.
func ProgressivelyReadFloatArray(in chan float32, out chan float32) {
	const bufferSize = 1024
	buffer := make([]float32, bufferSize)
	var v, n float64
	idx := 0

	for smpl := range in {
		buffer[idx] = smpl
		idx++

		// Process the buffer when it's full
		if idx == bufferSize {
			for _, val := range buffer {
				z := math.Abs(float64(val))
				if z > v {
					v += (z - v) / 8
				} else {
					v -= (v - z) / 512
				}

				n++
				if n == float64(utils.INTERVAL) {
					n = 0
					out <- float32(v)
				}
			}
			idx = 0
		}
	}

	// Process remaining samples
	for i := 0; i < idx; i++ {
		z := math.Abs(float64(buffer[i]))
		if z > v {
			v += (z - v) / 8
		} else {
			v -= (v - z) / 512
		}

		n++
		if n == float64(utils.INTERVAL) {
			n = 0
			out <- float32(v)
		}
	}

	close(out)
}
