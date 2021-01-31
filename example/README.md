Applying Digital Filter
=======================

This example does applying the band-pass filter to the audio signal.

## Required Tool

- [go](https://golang.org/dl/)
- [sox](http://sox.sourceforge.net)

## Usage

Open the terminal app and run the following steps:

```console
$ sox music.wav -t f64 input.raw
$ go run main.go
$ play -c 2 -r 44100 -t f64 output.raw
```

You hear the music that the band-pass filter applied.

NOTE: the sample rate of the `music.wav` must be 44100 Hz and the channels must be stereo.
