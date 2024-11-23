// /home/megalith/megasound/keydetector/keydetector.go
package keydetector

import (
	"errors"
	"github.com/rickcollette/megasound"
	"math"
)

type KeyDetector struct {
	Streamer   megasound.Streamer
	SampleRate int
}

func NewKeyDetector(s megasound.Streamer, sampleRate int) *KeyDetector {
	return &KeyDetector{Streamer: s, SampleRate: sampleRate}
}

func dynamicThreshold(pcp []float64) float64 {
	if len(pcp) == 0 {
		return 0.1
	}

	totalEnergy, mean, variance := 0.0, 0.0, 0.0

	for _, value := range pcp {
		totalEnergy += value
		mean += value
	}

	mean /= float64(len(pcp))

	for _, value := range pcp {
		variance += math.Pow(value-mean, 2)
	}

	variance /= float64(len(pcp))
	normalizedVariance := variance / totalEnergy

	baseThreshold := 0.1
	maxThreshold := 0.5
	threshold := baseThreshold + (normalizedVariance * maxThreshold)

	if threshold > maxThreshold {
		return maxThreshold
	} else if threshold < baseThreshold {
		return baseThreshold
	}

	return threshold
}

func (kd *KeyDetector) DetectKey() (KeyResult, error) {
	if kd.Streamer == nil {
		return KeyResult{}, errors.New("streamer cannot be nil")
	}

	pcp, err := ComputePCP(kd.Streamer, kd.SampleRate)
	if err != nil {
		return KeyResult{}, err
	}

	result := KrumhanslSchmuckler(pcp)
	if result.Confidence < dynamicThreshold(pcp) {
		return KeyResult{}, errors.New("low confidence in key detection")
	}

	return result, nil
}
