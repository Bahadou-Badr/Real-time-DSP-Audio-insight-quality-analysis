# Audio Insight & Quality Analysis Service

A backend service that analyzes digital audio files and extracts technical
and perceptual quality metrics such as loudness, dynamic range, spectral
distribution, and noise characteristics.

The system is designed with a small but real-world architecture inspired by
streaming and telecommunications platforms.

## Architecture (High Level)

- Client uploads audio via HTTP
- Go API manages jobs and lifecycle
- Python DSP worker performs audio analysis
- FFmpeg is used for decoding and normalization
- WebSocket streams analysis progress
- gRPC connects Go services to Python workers

## Tech Stack

- Go (HTTP API, WebSocket, orchestration)
- Python (DSP analysis)
- FFmpeg (audio processing)
- gRPC (internal communication)

-------------
### Offline analysis
Response Result

```bash
curl -X POST http://localhost:8080/analyze   -F "audio=@percLoop.wav"
```
```json
{
"id":"07926996-5999-43ea-b6b0-ca652cdbac27",
"status":"done",
"result":
    {
    "duration_seconds":7.5,
    "rms":0.01992344219382106,
    "peak":0.602294921875,
    "silence_ratio":0.6898503401360544,
    "spectral_centroid":4741.869914131183,
    "low_energy":31912.591789650374,
    "mid_energy":254045.28764874054,
    "high_energy":574943.9057751946
    },
"created_at":"2026-01-29T02:16:14.5216644+01:00",
"updated_at":"2026-01-29T02:16:14.6683508+01:00"
}
```
Offline DSP is the foundation of real-time DSP.