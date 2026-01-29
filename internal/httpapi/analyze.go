package httpapi

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"audio-insight-quality-analysis/dsp"
	"audio-insight-quality-analysis/internal/job"
)

type Handler struct {
	Jobs *job.Store
}

func (h *Handler) Analyze(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(50 << 20)

	file, _, err := r.FormFile("audio")
	if err != nil {
		http.Error(w, "audio file required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	tmp, err := os.CreateTemp("", "upload-*")
	if err != nil {
		http.Error(w, "temp file error", 500)
		return
	}
	defer os.Remove(tmp.Name())

	io.Copy(tmp, file)

	// ---- Create Job ----
	j := h.Jobs.Create()
	j.Status = job.StatusProcessing
	h.Jobs.Update(j)

	// ---- Decode + Analyze ----
	pcm, err := dsp.DecodeToPCM(tmp.Name())
	if err != nil {
		j.Status = job.StatusFailed
		j.Error = err.Error()
		h.Jobs.Update(j)
		json.NewEncoder(w).Encode(j)
		return
	}

	result := dsp.AnalyzeAudio(pcm.Samples, pcm.SampleRate)

	j.Status = job.StatusDone
	j.Result = result
	h.Jobs.Update(j)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(j)
}
