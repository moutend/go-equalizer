go-equalizer
============

`go-equalizer` provides equalizers based on the Robert Bristow-Johnson's audio EQ cookbook.

- [Cookbook formulae for audio EQ biquad filter coefficients](https://webaudio.github.io/Audio-EQ-Cookbook/audio-eq-cookbook.html)

This package supports the following digital filters:

- Low-pass
- High-pass
- All-pass
- Band-pass
- Band-reject
- Low-shelf
- High-shelf
- Peaking

## Install

```console
$ go get -u github.com/moutend/go-equalizer
```

## Example

See the `example` directory.

## Usage

```go
package main

import "github.com/moutend/go-equalizer/pkg/equalizer"

func main() {
	// Audio signal
	input := []float64{ /* ... */ }
	output := make([]float64, len(input))

	// Create the band-pass filter.
	bpf := equalizer.NewBandPass(44100, 440, 0.5)

	for i := range input {
		output[i] = bpf.Apply(input)
	}
}
```

NOTE: `go-equalizer` does not provide the way to read the audio file as a float64 slice.

## LICENSE

MIT

## Author

[Yoshiyuki Koyanagi <moutend@gmail.com>](https://github.com/moutend/)
