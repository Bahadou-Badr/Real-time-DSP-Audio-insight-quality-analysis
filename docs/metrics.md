## What is sample rate?
Sample rate defines how many times per second the audio signal is measured.
Higher sample rates capture more detail, but 44.1kHz is sufficient for most music and speech analysis.
## What is bit depth?
bit depth refers to the number of bits used to represent the amplitude (volume) of each audio sample in a digital signal. While sampling rate determines the frequency range, bit depth determines the precision of each "snapshot," directly impacting the signal's dynamic range and noise floor.
## Why PCM for analysis?
PCM is the preferred format for analysis because it is uncompressed and linear. Unlike lossy formats (MP3, AAC) that discard data to save space, PCM preserves the original mathematical relationship between samples.
## Why normalization before DSP?
Many DSP effects especially non-linear ones like compressors, saturators, or certain filters have "thresholds." Normalizing ensures your input signal consistently hits these thresholds, making your Spectral Analysis or processing results repeatable across different files.

**Peak vs. Loudness (LUFS) Normalization**

- **Peak Normalization:** Adjusts based on the single highest sample. This is standard for technical **PCM** analysis.
- **Loudness (LUFS) Normalization:** Adjusts based on *perceived* average volume. This is the standard for streaming platforms like Spotify or Apple Music to ensure listeners don't have to keep reaching for the volume knob.