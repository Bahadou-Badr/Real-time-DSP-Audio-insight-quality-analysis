package dsp

import (
	"math/cmplx"

	"gonum.org/v1/gonum/dsp/fourier"
)

func ComputeSpectrum(samples []float64, sampleRate int) (
	centroid float64,
	low float64,
	mid float64,
	high float64,
) {
	n := len(samples)
	if n == 0 {
		return
	}

	// Real FFT (this expects []float64)
	fft := fourier.NewFFT(n)
	freq := fft.Coefficients(nil, samples)

	// Frequency resolution
	binHz := float64(sampleRate) / float64(n)

	var magSum, weightedSum float64

	lowMax := 200.0
	midMax := 2000.0

	// Analyze positive frequencies
	for i := 1; i < len(freq); i++ {
		mag := cmplx.Abs(freq[i])
		freqHz := float64(i) * binHz

		magSum += mag
		weightedSum += freqHz * mag

		switch {
		case freqHz <= lowMax:
			low += mag
		case freqHz <= midMax:
			mid += mag
		default:
			high += mag
		}
	}

	if magSum > 0 {
		centroid = weightedSum / magSum
	}

	return
}
