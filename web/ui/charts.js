const MAX_POINTS = 200;

function createSeries() {
  return [];
}

function drawChart(canvas, data, color, maxValue) {
  const ctx = canvas.getContext("2d");
  const w = canvas.width;
  const h = canvas.height;

  ctx.clearRect(0, 0, w, h);

  ctx.strokeStyle = color;
  ctx.beginPath();

  data.forEach((v, i) => {
    const x = (i / MAX_POINTS) * w;
    const y = h - (v / maxValue) * h;
    if (i === 0) ctx.moveTo(x, y);
    else ctx.lineTo(x, y);
  });

  ctx.stroke();
}

export class Charts {
  constructor() {
    this.rms = createSeries();
    this.centroid = createSeries();

    this.rmsCanvas = document.getElementById("rmsChart");
    this.centroidCanvas = document.getElementById("centroidChart");
  }

  push(stats) {
    this.rms.push(stats.rms);
    this.centroid.push(stats.spectral_centroid);

    if (this.rms.length > MAX_POINTS) this.rms.shift();
    if (this.centroid.length > MAX_POINTS) this.centroid.shift();

    drawChart(this.rmsCanvas, this.rms, "lime", 0.05);
    drawChart(this.centroidCanvas, this.centroid, "cyan", 8000);
  }
}
