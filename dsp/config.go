package dsp

const (
	FrameSize  = 1024
	HopSize    = 512
	SampleRate = 44100

	// ~1 second window: 44100 / 512 ≈ 86 frames
	StreamWindowFrames = 80
)
