// /home/megalith/megasound/keydetector/pcp.go
package keydetector

import (
	"errors"
	"math"
	"math/cmplx"

	"gonum.org/v1/gonum/dsp/fourier"
	"gonum.org/v1/gonum/dsp/window"

	"github.com/rickcollette/megasound"
)

// ComputePCP analyzes an audio stream and returns the detected key with Camelot Wheel notation.
func ComputePCP(streamer megasound.Streamer, sampleRate int) ([]float64, error) {
	pcp := make([]float64, 12)

	samples := make([][2]float64, 1024)
	mono := make([]float64, 1024)

	for {
		n, ok := streamer.Stream(samples)
		if n == 0 || !ok {
			break
		}

		frequency := estimateFrequency(samples[:n], sampleRate, mono)
		if frequency > 0 {
			pitchClass := computePitchClass(frequency)
			pcp[pitchClass]++
		}
	}

	total := 0.0
	for _, v := range pcp {
		total += v
	}
	if total == 0 {
		return nil, errors.New("no valid PCP data detected")
	}
	for i := range pcp {
		pcp[i] /= total
	}

	return pcp, nil
}

// computePitchClass calculates the pitch class (0-11) for a given frequency.
func computePitchClass(frequency float64) int {
	note := math.Mod(math.Log2(frequency/440.0)*12+69, 12)
	if note < 0 {
		note += 12
	}
	return int(note)
}

// estimateFrequency estimates the fundamental frequency of a given audio sample using FFT.
func estimateFrequency(samples [][2]float64, sampleRate int, mono []float64) float64 {
	for i, s := range samples {
		mono[i] = (s[0] + s[1]) / 2
	}

	window.Hann(mono)

	fft := fourier.NewFFT(len(mono))
	complexSpectrum := fft.Coefficients(nil, mono)

	magnitudes := make([]float64, len(complexSpectrum))
	for i, c := range complexSpectrum {
		magnitudes[i] = cmplx.Abs(c)
	}

	peakIndex := 0
	for i, mag := range magnitudes[:len(magnitudes)/2] {
		if mag > magnitudes[peakIndex] {
			peakIndex = i
		}
	}

	return float64(peakIndex) * float64(sampleRate) / float64(len(mono))
}
