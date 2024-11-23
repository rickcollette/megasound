package keydetector

// KeyProfile represents a key profile with its precomputed norm.
type KeyProfile struct {
	Profile []float64
	Norm    float64
}

// PCPKeyProfiles returns predefined key profiles for all major and minor keys
// based on music theory principles. These profiles correspond to the harmonic
// relationships between the 12 notes in the chromatic scale for each key.
func PCPKeyProfiles() map[string][]float64 {
	// Major key profiles based on the harmonic distribution of intervals in the major scale
	return map[string][]float64{
		// Major keys (C Major, G Major, etc.)
		"C Major":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"G Major":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"D Major":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"A Major":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"E Major":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"B Major":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"F# Major": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"Db Major": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"Ab Major": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"Eb Major": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},
		"Bb Major": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84},

		// Minor keys (A Minor, E Minor, etc.)
		"A Minor":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"E Minor":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"B Minor":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"F# Minor": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"Db Minor": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"Ab Minor": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"Eb Minor": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"Bb Minor": {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"F Minor":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"C Minor":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"G Minor":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
		"D Minor":  {1.0, 0.88, 0.84, 0.88, 1.0, 0.84, 0.88, 0.84, 0.88, 0.88, 0.84, 0.88},
	}
}

func KrumhanslKeyProfiles() map[string][]float64 {
	return map[string][]float64{
		"C Major":    {6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19, 2.39, 3.66, 2.29, 2.88},
		"G Major":    {5.19, 2.39, 3.66, 2.29, 2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52},
		"D Major":    {4.38, 4.09, 2.52, 5.19, 2.39, 3.66, 2.29, 2.88, 6.35, 2.23, 3.48, 2.33},
		"A Major":    {2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19, 2.39, 3.66, 2.29},
		"E Major":    {2.29, 2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19, 2.39, 3.66},
		"B Major":    {3.66, 2.29, 2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19, 2.39},
		"F# Major":   {2.39, 3.66, 2.29, 2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19},
		"Db Major":   {5.19, 2.39, 3.66, 2.29, 2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52},
		"Ab Major":   {4.38, 4.09, 2.52, 5.19, 2.39, 3.66, 2.29, 2.88, 6.35, 2.23, 3.48, 2.33},
		"Eb Major":   {2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19, 2.39, 3.66, 2.29},
		"Bb Major":   {2.29, 2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19, 2.39, 3.66},
		"F Major":    {3.66, 2.29, 2.88, 6.35, 2.23, 3.48, 2.33, 4.38, 4.09, 2.52, 5.19, 2.39},
		"A Minor":    {6.33, 2.68, 3.52, 5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96, 4.52, 2.98},
		"E Minor":    {5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96, 4.52, 2.98, 3.30, 2.27, 2.99},
		"B Minor":    {3.92, 3.53, 2.96, 4.52, 2.98, 3.30, 2.27, 2.99, 5.38, 2.60, 3.53, 2.54},
		"F# Minor":   {4.52, 2.98, 3.30, 2.27, 2.99, 5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96},
		"C# Minor":   {2.99, 5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96, 4.52, 2.98, 3.30, 2.27},
		"G# Minor":   {2.27, 2.99, 5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96, 4.52, 2.98, 3.30},
		"D# Minor":   {3.30, 2.27, 2.99, 5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96, 4.52, 2.98},
		"Bb Minor":   {4.52, 2.98, 3.30, 2.27, 2.99, 5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96},
		"F Minor":    {3.92, 3.53, 2.96, 4.52, 2.98, 3.30, 2.27, 2.99, 5.38, 2.60, 3.53, 2.54},
		"C Minor":    {2.96, 4.52, 2.98, 3.30, 2.27, 2.99, 5.38, 2.60, 3.53, 2.54, 3.92, 3.53},
		"G Minor":    {2.98, 3.30, 2.27, 2.99, 5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96, 4.52},
		"D Minor":    {5.38, 2.60, 3.53, 2.54, 3.92, 3.53, 2.96, 4.52, 2.98, 3.30, 2.27, 2.99},
	}
}