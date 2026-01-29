package dsp

func AnalyzeAudio(samples []float64, sampleRate int) AnalysisResult {
	// Duration
	duration := float64(len(samples)) / float64(sampleRate)

	// RMS + Peak
	rms, peak := ComputeRMSAndPeak(samples)

	// Silence
	silenceRatio := ComputeSilenceRatio(samples, 0.001)

	// FFT-based metrics (single call, efficient)
	centroid, low, mid, high := ComputeSpectrum(samples, sampleRate)

	return AnalysisResult{
		DurationSeconds:  duration,
		RMS:              rms,
		Peak:             peak,
		SilenceRatio:     silenceRatio,
		SpectralCentroid: centroid,
		LowEnergy:        low,
		MidEnergy:        mid,
		HighEnergy:       high,
	}
}
