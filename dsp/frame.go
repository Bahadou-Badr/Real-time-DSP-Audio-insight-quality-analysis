package dsp

func AnalyzeFrame(frame []float64, sampleRate int) FrameResult {
	rms, peak := ComputeRMSAndPeak(frame)
	silence := rms < 0.001

	centroid, low, mid, high := ComputeSpectrum(frame, sampleRate)

	return FrameResult{
		RMS:              rms,
		Peak:             peak,
		Silence:          silence,
		SpectralCentroid: centroid,
		LowEnergy:        low,
		MidEnergy:        mid,
		HighEnergy:       high,
	}
}
