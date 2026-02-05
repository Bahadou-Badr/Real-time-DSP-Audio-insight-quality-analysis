package httpapi

import (
	"encoding/json"
	"net/http"

	"audio-insight-quality-analysis/dsp"
)

func (h *Handler) AnalyzeStreamFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		http.Error(w, "missing path", 400)
		return
	}

	reader, err := dsp.NewPCMFileStream(path, 1024, true)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	result := dsp.AnalyzeStream(reader)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
