package dsp

import "io"

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

/* -- STREAMING  */
func AnalyzeStream(reader *PCMStreamReader) AnalysisResult {
	var (
		buffer []float64
		stats  StreamStats
		total  int
	)

	for {
		chunk, err := reader.ReadSamples()
		if len(chunk) > 0 {
			buffer = append(buffer, chunk...)
			total += len(chunk)
		}

		for len(buffer) >= FrameSize {
			frame := buffer[:FrameSize]
			buffer = buffer[HopSize:]

			fr := AnalyzeFrame(frame, reader.SampleRate)
			stats.Push(fr)
		}

		if err == io.EOF {
			break
		}
	}

	result := stats.Finalize()
	result.DurationSeconds = float64(total) / float64(reader.SampleRate)
	return result
}
