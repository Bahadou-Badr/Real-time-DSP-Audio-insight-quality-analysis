package dsp

type AnalysisResult struct {
	DurationSeconds  float64 `json:"duration_seconds"`
	RMS              float64 `json:"rms"`
	Peak             float64 `json:"peak"`
	SilenceRatio     float64 `json:"silence_ratio"`
	SpectralCentroid float64 `json:"spectral_centroid"`
	LowEnergy        float64 `json:"low_energy"`
	MidEnergy        float64 `json:"mid_energy"`
	HighEnergy       float64 `json:"high_energy"`
}

type FrameResult struct {
	RMS              float64
	Peak             float64
	Silence          bool
	SpectralCentroid float64
	LowEnergy        float64
	MidEnergy        float64
	HighEnergy       float64
}
