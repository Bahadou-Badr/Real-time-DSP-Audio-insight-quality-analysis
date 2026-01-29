package dsp

func ComputeSilenceRatio(samples []float64, threshold float64) float64 {
	var silent int
	for _, s := range samples {
		abs := s
		if abs < 0 {
			abs = -abs
		}
		if abs < threshold {
			silent++
		}
	}
	return float64(silent) / float64(len(samples))
}
