export function connectWS(onStats) {
  const ws = new WebSocket("ws://localhost:8080/ws/audio");
  ws.binaryType = "arraybuffer";

  const client = {
    sendAudio(buffer) {
      if (ws.readyState === WebSocket.OPEN) {
        ws.send(buffer);
      }
    }
  };

  ws.onmessage = (evt) => {
    const data = JSON.parse(evt.data);
    onStats(data);
  };

  return client;
}
