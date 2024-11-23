// /home/megalith/megasound/keydetector/krumhansl.go
package keydetector

import (
	"math"
)

// KeyResult represents the detected key with additional information like Camelot Wheel notation.
type KeyResult struct {
	Key             string
	CamelotNotation string
	Confidence      float64
}

// KrumhanslSchmuckler computes the most likely key based on PCP and includes Camelot Wheel notation.
func KrumhanslSchmuckler(pcp []float64) KeyResult {
	profiles := KrumhanslKeyProfiles()
	camelotMapping := camelotWheelMapping()

	var bestKey string
	var bestScore float64

	for key, profile := range profiles {
		score := cosineSimilarity(pcp, profile)
		if score > bestScore {
			bestScore = score
			bestKey = key
		}
	}

	return KeyResult{
		Key:             bestKey,
		CamelotNotation: camelotMapping[bestKey],
		Confidence:      bestScore,
	}
}

// cosineSimilarity computes the cosine similarity between two vectors.
func cosineSimilarity(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0
	}

	dotProduct := 0.0
	normA, normB := 0.0, 0.0

	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// camelotWheelMapping returns a map of keys to their Camelot Wheel notations.
func camelotWheelMapping() map[string]string {
	return map[string]string{
		"A minor": "1A", "C major": "8A",
		"E minor": "2A", "D major": "2B",
		"B minor": "3A", "A major": "3B",
		"F# minor": "4A", "B major": "4B",
		"C# minor": "5A", "E major": "5B",
		"G# minor": "6A", "D# minor": "7A",
		"F minor": "8A", "G major": "7B",
		"B♭ minor": "9A", "D♭ major": "9B",
		"G minor": "10A", "F major": "10B",
		"D minor": "11A", "E♭ major": "11B",
		"A♭ minor": "12A", "C minor": "1B",
	}
}
