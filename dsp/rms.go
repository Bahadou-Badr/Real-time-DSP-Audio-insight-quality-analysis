package dsp

import "math"

func ComputeRMSAndPeak(samples []float64) (rms, peak float64) {
	var sumSquares float64

	for _, s := range samples {
		sumSquares += s * s
		if abs := math.Abs(s); abs > peak {
			peak = abs
		}
	}

	rms = math.Sqrt(sumSquares / float64(len(samples)))
	return
}
