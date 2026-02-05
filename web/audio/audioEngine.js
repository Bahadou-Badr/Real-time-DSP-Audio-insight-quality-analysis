let audioCtx = null;
let processor = null;
let stream = null;

export async function startAudio(wsClient) {
  stream = await navigator.mediaDevices.getUserMedia({ audio: true });

  audioCtx = new AudioContext({ sampleRate: 48000 });
  await audioCtx.audioWorklet.addModule("./audio/audioProcessor.js");

  const source = audioCtx.createMediaStreamSource(stream);
  processor = new AudioWorkletNode(audioCtx, "audio-processor");

  processor.port.onmessage = (e) => {
    wsClient.sendAudio(e.data);
  };

  source.connect(processor);
}

export function stopAudio() {
  if (processor) {
    processor.disconnect();
    processor = null;
  }

  if (audioCtx) {
    audioCtx.close();
    audioCtx = null;
  }

  if (stream) {
    stream.getTracks().forEach(t => t.stop());
    stream = null;
  }
}
