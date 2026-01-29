#!/bin/bash

# Convert any input audio to a normalized internal WAV format
ffmpeg -y \
  -i "$1" \
  -ar 44100 \
  -ac 2 \
  -sample_fmt s16 \
  output.wav