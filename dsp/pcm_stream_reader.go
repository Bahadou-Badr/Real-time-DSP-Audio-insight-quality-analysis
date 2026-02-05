package dsp

import (
	"encoding/binary"
	"io"
	"os"
	"time"
)

type PCMStreamReader struct {
	r          io.Reader
	SampleRate int
	ChunkSize  int // samples per read
	RealTime   bool
}

func NewPCMFileStream(path string, chunkSize int, realTime bool) (*PCMStreamReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &PCMStreamReader{
		r:          f,
		SampleRate: SampleRate,
		ChunkSize:  chunkSize,
		RealTime:   realTime,
	}, nil
}

func (p *PCMStreamReader) ReadSamples() ([]float64, error) {
	buf := make([]byte, p.ChunkSize*2) // int16
	n, err := p.r.Read(buf)
	if n == 0 {
		return nil, err
	}

	samples := make([]float64, n/2)
	for i := 0; i < len(samples); i++ {
		v := int16(binary.LittleEndian.Uint16(buf[i*2:]))
		samples[i] = float64(v) / 32768.0
	}

	if p.RealTime {
		d := time.Duration(float64(len(samples))/float64(p.SampleRate)*1000) * time.Millisecond
		time.Sleep(d)
	}

	return samples, err
}
