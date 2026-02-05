import { startAudio, stopAudio } from "./audio/audioEngine.js";
import { connectWS } from "./ws/wsClient.js";
import { updateMeters } from "./ui/meters.js";
import { Charts } from "./ui/charts.js";

let wsClient = null;
let charts = null;
let running = false;
let lastStats = null;

const startBtn = document.getElementById("startBtn");
const stopBtn = document.getElementById("stopBtn");

startBtn.onclick = async () => {
  if (running) return;

  console.log("Start clicked");
  running = true;

  charts = new Charts();

  wsClient = connectWS((stats) => {
    if (!running) return;

    lastStats = stats;
    updateMeters(stats);
    charts.push(stats);
  });

  await startAudio(wsClient);

  startBtn.disabled = true;
  stopBtn.disabled = false;
};

stopBtn.onclick = () => {
  if (!running) return;

  console.log("Stop clicked");
  running = false;

  stopAudio();

  startBtn.disabled = false;
  stopBtn.disabled = true;

  if (lastStats) {
    // Freeze meters
    updateMeters(lastStats);

    // Render final result in UI
    const out = document.getElementById("finalResult");
    out.textContent = JSON.stringify(
      {
        duration_seconds: lastStats.duration_seconds.toFixed(2),
        rms: lastStats.rms,
        peak: lastStats.peak,
        silence_ratio: lastStats.silence_ratio,
        spectral_centroid: lastStats.spectral_centroid,
        low_energy: lastStats.low_energy,
        mid_energy: lastStats.mid_energy,
        high_energy: lastStats.high_energy
      },
      null,
      2
    );
  }
};

