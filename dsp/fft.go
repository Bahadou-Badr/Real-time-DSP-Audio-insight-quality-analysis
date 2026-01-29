package dsp

import "gonum.org/v1/gonum/dsp/fourier"

func FFT(samples []float64) []complex128 {
	fft := fourier.NewFFT(len(samples))
	return fft.Coefficients(nil, samples)
}
