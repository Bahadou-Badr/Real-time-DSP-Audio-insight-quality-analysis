package main

import (
	"log"
	"net/http"

	"audio-insight-quality-analysis/internal/httpapi"
	"audio-insight-quality-analysis/internal/job"
)

func main() {
	jobStore := job.NewStore()
	handler := &httpapi.Handler{Jobs: jobStore}

	http.HandleFunc("/analyze", handler.Analyze)

	log.Println("API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
