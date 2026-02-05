export function updateMeters(stats) {
  // RMS & Peak
  document.getElementById("rms").innerText =
    "RMS: " + stats.rms.toFixed(6);

  document.getElementById("peak").innerText =
    "Peak: " + stats.peak.toFixed(6);

  document.getElementById("lowEnergyText").innerText =
    "Low: " + stats.low_energy.toFixed(4);

  document.getElementById("midEnergyText").innerText =
    "Mid: " + stats.mid_energy.toFixed(4);

  document.getElementById("highEnergyText").innerText =
    "High: " + stats.high_energy.toFixed(4);
}
