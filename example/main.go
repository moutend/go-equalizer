package main

import (
	"encoding/binary"
	"io/ioutil"
	"math"

	"github.com/moutend/go-equalizer/pkg/equalizer"
)

func main() {
	data, err := ioutil.ReadFile("input.raw")

	if err != nil {
		panic(err)
	}

	// f0 -> L channel / f1 -> R channel
	f0 := equalizer.NewBandPass(44100, 440, 0.5)
	f1 := equalizer.NewBandPass(44100, 440, 0.5)

	ch := 0
	bs := []byte{}

	for i := 0; i < len(data); i += 8 {
		input := math.Float64frombits(
			binary.LittleEndian.Uint64(data[i : i+8]),
		)

		output := input

		if ch == 0 {
			output = f0.Apply(output)
		} else {
			output = f1.Apply(output)
		}

		ch = (ch + 1) % 2

		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, math.Float64bits(output))

		bs = append(bs, b...)
	}
	if err := ioutil.WriteFile("output.raw", bs, 0644); err != nil {
		panic(err)
	}
}
