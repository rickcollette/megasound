package detection

import (
	"github.com/rickcollette/megasound/bpm/utils"
	"math"
)

// CalcChunkLen calculates the chunk size for a given duration in seconds.
func CalcChunkLen(seconds int) int {
	return (utils.RATE / utils.INTERVAL) * seconds
}

func ScanForBpm(nrg []float32, slowest, fastest float64, steps, samples int) float64 {
	coarseStep := (bpmToInterval(slowest) - bpmToInterval(fastest)) / float64(steps)
	refinedStep := coarseStep / 10

	bestInterval := float64(0)
	bestScore := math.Inf(1)

	for interval := bpmToInterval(fastest); interval <= bpmToInterval(slowest); interval += coarseStep {
		score := 0.0
		for s := 0; s < samples; s++ {
			score += autodifference(nrg, interval) // Use autodifference here
		}
		if score < bestScore {
			bestScore = score
			bestInterval = interval
		}
	}

	// Refine search around the best interval
	for interval := bestInterval - refinedStep; interval <= bestInterval+refinedStep; interval += refinedStep {
		score := 0.0
		for s := 0; s < samples; s++ {
			score += autodifference(nrg, interval) // Use autodifference here
		}
		if score < bestScore {
			bestScore = score
			bestInterval = interval
		}
	}

	return intervalToBpm(bestInterval) // Use intervalToBpm here
}
