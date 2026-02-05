package main

import (
	"log"
	"net/http"

	"audio-insight-quality-analysis/internal/httpapi"
	"audio-insight-quality-analysis/internal/job"
	"audio-insight-quality-analysis/internal/ws"
)

// inside main()

func main() {
	jobStore := job.NewStore()
	handler := &httpapi.Handler{Jobs: jobStore}

	http.HandleFunc("/analyze", handler.Analyze)

	http.HandleFunc("/analyze/stream-file", handler.AnalyzeStreamFile)
	http.HandleFunc("/ws/audio", ws.AudioStreamHandler)
	log.Println("API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
