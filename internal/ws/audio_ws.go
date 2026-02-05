package ws

import (
	"encoding/binary"
	"log"
	"math"
	"net/http"

	"audio-insight-quality-analysis/dsp"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func AudioStreamHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	stats := dsp.NewStreamStats()
	buffer := make([]float64, 0, dsp.FrameSize*2)
	totalSamples := 0

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("ws closed:", err)
			break
		}

		// float32 PCM → float64
		for i := 0; i+4 <= len(msg); i += 4 {
			v := binary.LittleEndian.Uint32(msg[i:])
			f := float64(math.Float32frombits(v))
			buffer = append(buffer, f)
			totalSamples++ // counting each sample once
		}

		// process frames
		for len(buffer) >= dsp.FrameSize {
			frame := buffer[:dsp.FrameSize]
			buffer = buffer[dsp.HopSize:] // slide buffer
			fr := dsp.AnalyzeFrame(frame, dsp.SampleRate)
			stats.Push(fr)

			// compute live duration
			partial := stats.Finalize()
			partial.DurationSeconds = float64(totalSamples) / float64(dsp.SampleRate)

			conn.WriteJSON(partial)
		}

	}

	final := stats.Finalize()
	// final.DurationSeconds = float64(totalSamples) / float64(dsp.SampleRate)
	conn.WriteJSON(final)
}
