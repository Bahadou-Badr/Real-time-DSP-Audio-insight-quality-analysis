package dsp

type StreamStats struct {
	frames []FrameResult
}

func NewStreamStats() *StreamStats {
	return &StreamStats{
		frames: make([]FrameResult, 0, StreamWindowFrames),
	}
}

func (s *StreamStats) Push(f FrameResult) {
	// append new frame
	s.frames = append(s.frames, f)

	// trim old frames
	if len(s.frames) > StreamWindowFrames {
		s.frames = s.frames[1:]
	}
}

func (s *StreamStats) Finalize() AnalysisResult {
	if len(s.frames) == 0 {
		return AnalysisResult{}
	}

	var (
		rmsSum      float64
		peak        float64
		silentCount int
		centroidSum float64
		low         float64
		mid         float64
		high        float64
	)

	for _, f := range s.frames {
		rmsSum += f.RMS
		centroidSum += f.SpectralCentroid
		low += f.LowEnergy
		mid += f.MidEnergy
		high += f.HighEnergy

		if f.Peak > peak {
			peak = f.Peak
		}
		if f.Silence {
			silentCount++
		}
	}

	n := float64(len(s.frames))

	return AnalysisResult{
		RMS:              rmsSum / n,
		Peak:             peak,
		SilenceRatio:     float64(silentCount) / n,
		SpectralCentroid: centroidSum / n,
		LowEnergy:        low / n,
		MidEnergy:        mid / n,
		HighEnergy:       high / n,
	}
}
