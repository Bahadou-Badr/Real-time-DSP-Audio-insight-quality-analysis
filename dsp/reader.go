package dsp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os/exec"
)

type PCMData struct {
	Samples    []float64
	SampleRate int
}

// Decode audio file → mono PCM float64 using FFmpeg
func DecodeToPCM(path string) (*PCMData, error) {
	// FFmpeg command:
	// - convert to mono
	// - 16-bit signed little endian
	// - raw output to stdout
	cmd := exec.Command(
		"ffmpeg",
		"-i", path,
		"-ac", "1",
		"-ar", "44100",
		"-f", "s16le",
		"-hide_banner",
		"-loglevel", "error",
		"pipe:1",
	)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("ffmpeg decode failed: %w", err)
	}

	raw := out.Bytes()
	samples := make([]float64, len(raw)/2)

	// Convert int16 → float64 (-1.0 to 1.0)
	for i := 0; i < len(samples); i++ {
		v := int16(binary.LittleEndian.Uint16(raw[i*2:]))
		samples[i] = float64(v) / 32768.0
	}

	return &PCMData{
		Samples:    samples,
		SampleRate: SampleRate,
	}, nil
}

/* STREAMING */
func NewPCMStreamReader(path string) (*PCMStreamReader, error) {
	cmd := exec.Command(
		"ffmpeg",
		"-i", path,
		"-ac", "1",
		"-ar", "44100",
		"-f", "s16le",
		"-hide_banner",
		"-loglevel", "error",
		"pipe:1",
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return &PCMStreamReader{
		r:          stdout,
		SampleRate: SampleRate,
	}, nil
}

// Read next chunk of float64 samples
// func (p *PCMStreamReader) ReadSamples() ([]float64, error) {
// 	buf := make([]byte, 4096)
// 	n, err := p.r.Read(buf)
// 	if n == 0 {
// 		return nil, err
// 	}

// 	samples := make([]float64, 0, n/2)
// 	for i := 0; i+1 < n; i += 2 {
// 		v := int16(binary.LittleEndian.Uint16(buf[i:]))
// 		samples = append(samples, float64(v)/32768.0)
// 	}

// 	return samples, err
// }
