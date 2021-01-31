// Package equalizer provides equalizers based on the Robert Bristow-Johnson's audio EQ cookbook.
//
// This package supports the following digital filters:
//
//     - Low-pass
//     - High-pass
//     - All-pass
//     - Band-pass
//     - Band-reject
//     - Low-shelf
//     - High-shelf
//     - Peaking
package equalizer

import "math"

// Mode represents the kind of digital filters.
type FilterName int

// FilterName constants are digital filter names.
const (
	Undefined FilterName = iota
	LowPass
	HighPass
	AllPass
	BandPass
	BandReject
	LowShelf
	HighShelf
	Peaking
)

// Pi value is used as the default pi value in this package.
const Pi = 3.1415926535897932384626433

var (
	p = Pi
)

// SetPi sets the pi value. After calling this function, call the constructor function such as NewLowPass().
func SetPi(value float64) {
	p = value
}

// UnsetPi sets the pi value to default value.
func UnsetPi() {
	p = Pi
}

// Filter holds the digital filter parameters.
type Filter struct {
	name FilterName

	// state variables
	in1  float64
	in2  float64
	out1 float64
	out2 float64

	// digital filter parameters
	a0 float64
	a1 float64
	a2 float64
	b0 float64
	b1 float64
	b2 float64
}

// IsZero returns true when the f is not initialized.
func (f *Filter) IsZero() bool {
	return f.name == Undefined
}

// FilterName returns filter name.
func (f *Filter) Name() FilterName {
	return f.name
}

// Apply applies the current filter and returns the value.
func (f *Filter) Apply(input float64) float64 {
	output := (f.b0/f.a0)*input +
		(f.b1/f.a0)*f.in1 +
		(f.b2/f.a0)*f.in2 -
		(f.a1/f.a0)*f.out1 -
		(f.a2/f.a0)*f.out2

	f.in2 = f.in1
	f.in1 = input

	f.out2 = f.out1
	f.out1 = output

	return output
}

// NewLowPass returns the low-pass filter.
//
// Parameters:
//
//     - sampleRate ... sample rate in Hz. e.g. 44100.0
//     - frequency ... Cut off frequency in Hz.
//     - q ... Q value.
//
// NOTE: q must be greater than 0.
func NewLowPass(sampleRate, frequency, q float64) *Filter {
	w0 := 2.0 * p * frequency / sampleRate
	alpha := math.Sin(w0) / (2.0 * q)

	return &Filter{
		name: LowPass,
		a0:   1.0 + alpha,
		a1:   -2.0 * math.Cos(w0),
		a2:   1.0 - alpha,
		b0:   (1.0 - math.Cos(w0)) / 2.0,
		b1:   1.0 - math.Cos(w0),
		b2:   (1.0 - math.Cos(w0)) / 2.0,
	}
}

// NewHighPass returns the high-pass filter.
//
// Parameters:
//
//     - sampleRate ... sample rate in Hz. e.g. 44100.0
//     - frequency ... Cut off frequency in Hz.
//     - q ... Q value.
//
// NOTE: q must be greater than 0.
func NewHighPass(sampleRate, frequency, q float64) *Filter {
	w0 := 2.0 * p * frequency / sampleRate
	alpha := math.Sin(w0) / (2.0 * q)

	return &Filter{
		name: HighPass,
		a0:   1.0 + alpha,
		a1:   -2.0 * math.Cos(w0),
		a2:   1.0 - alpha,
		b0:   (1.0 + math.Cos(w0)) / 2.0,
		b1:   -1.0 * (1.0 + math.Cos(w0)),
		b2:   (1.0 + math.Cos(w0)) / 2.0,
	}
}

// NewAllPass returns the all-pass filter.
//
// Parameters:
//
//     - sampleRate ... sample rate in Hz. e.g. 44100.0
//     - frequency ... Cut off frequency in Hz.
//     - q ... Q value.
//
// NOTE: q must be greater than 0.
func NewAllPass(sampleRate, frequency, q float64) *Filter {
	w0 := 2.0 * p * frequency / sampleRate
	alpha := math.Sin(w0) / (2.0 * q)

	return &Filter{
		name: AllPass,
		a0:   1.0 + alpha,
		a1:   -2.0 * math.Cos(w0),
		a2:   1.0 - alpha,
		b0:   1.0 - alpha,
		b1:   -2.0 * math.Cos(w0),
		b2:   1.0 + alpha,
	}
}

// NewBandPass returns the band-pass filter.
//
// Parameters:
//
//     - sampleRate ... sample rate in Hz. e.g. 44100.0
//     - frequency ... Cut off frequency in Hz.
//     - width ... Band width.
//
// NOTE: width must be greater than 0.
func NewBandPass(sampleRate, frequency, width float64) *Filter {
	w0 := 2.0 * p * frequency / sampleRate
	alpha := math.Sin(w0) * math.Sinh(math.Log(2.0)/2.0*width*w0/math.Sin(w0))

	return &Filter{
		name: BandPass,
		a0:   1.0 + alpha,
		a1:   -2.0 * math.Cos(w0),
		a2:   1.0 - alpha,
		b0:   alpha,
		b1:   0.0,
		b2:   -1.0 * alpha,
	}
}

// NewBandReject returns the band-reject filter.
//
// Parameters:
//
//     - sampleRate ... sample rate in Hz. e.g. 44100.0
//     - frequency ... Cut off frequency in Hz.
//     - width ... Band width.
//
// NOTE: width must be greater than 0.
func NewBandReject(sampleRate, frequency, width float64) *Filter {
	w0 := 2.0 * p * frequency / sampleRate
	alpha := math.Sin(w0) * math.Sinh(math.Log(2.0)/2.0*width*w0/math.Sin(w0))

	return &Filter{
		name: BandReject,
		a0:   1.0 + alpha,
		a1:   -2.0 * math.Cos(w0),
		a2:   1.0 - alpha,
		b0:   1.0,
		b1:   -2.0 * math.Cos(w0),
		b2:   1.0,
	}
}

// NewLowShelf returns the low-shelf filter.
//
// Parameters:
//
//     - sampleRate ... sample rate in Hz. e.g. 44100.0
//     - frequency ... Cut off frequency in Hz.
//     - q ... Q value.
//     - gain ... Gain value in dB.
//
// NOTE: q must be greater than 0.
func NewLowShelf(sampleRate, frequency, q, gain float64) *Filter {
	w0 := 2.0 * p * frequency / sampleRate
	a := math.Pow(10.0, (gain / 40.0))
	beta := math.Sqrt(a) / q

	return &Filter{
		name: LowShelf,
		a0:   (a + 1.0) + (a-1.0)*math.Cos(w0) + beta*math.Sin(w0),
		a1:   -2.0 * ((a - 1.0) + (a+1.0)*math.Cos(w0)),
		a2:   (a + 1.0) + (a-1.0)*math.Cos(w0) - beta*math.Sin(w0),
		b0:   a * ((a + 1.0) - (a-1.0)*math.Cos(w0) + beta*math.Sin(w0)),
		b1:   2.0 * a * ((a - 1.0) - (a+1.0)*math.Cos(w0)),
		b2:   a * ((a + 1.0) - (a-1.0)*math.Cos(w0) - beta*math.Sin(w0)),
	}
}

// NewHighShelf returns the high-shelf filter.
//
// Parameters:
//
//     - sampleRate ... sample rate in Hz. e.g. 44100.0
//     - frequency ... Cut off frequency in Hz.
//     - q ... Q value.
//     - gain ... Gain value in dB.
//
// NOTE: q must be greater than 0.
func NewHighShelf(sampleRate, frequency, q, gain float64) *Filter {
	w0 := 2.0 * p * frequency / sampleRate
	a := math.Pow(10.0, (gain / 40.0))
	beta := math.Sqrt(a) / q

	return &Filter{
		name: HighShelf,
		a0:   (a + 1.0) - (a-1.0)*math.Cos(w0) + beta*math.Sin(w0),
		a1:   2.0 * ((a - 1.0) - (a+1.0)*math.Cos(w0)),
		a2:   (a + 1.0) - (a-1.0)*math.Cos(w0) - beta*math.Sin(w0),
		b0:   a * ((a + 1.0) + (a-1.0)*math.Cos(w0) + beta*math.Sin(w0)),
		b1:   -2.0 * a * ((a - 1.0) + (a+1.0)*math.Cos(w0)),
		b2:   a * ((a + 1.0) + (a-1.0)*math.Cos(w0) - beta*math.Sin(w0)),
	}
}

// NewPeaking returns the peaking-shelf filter.
//
// Parameters:
//
//     - sampleRate ... sample rate in Hz. e.g. 44100.0
//     - frequency ... Cut off frequency in Hz.
//     - width ... Width value.
//     - gain ... Gain value in dB.
//
// NOTE: width must be greater than 0.
func NewPeaking(sampleRate, frequency, width, gain float64) *Filter {
	w0 := 2.0 * p * frequency / sampleRate
	alpha := math.Sin(w0) * math.Sinh(math.Log(2.0)/2.0*width*w0/math.Sin(w0))
	a := math.Pow(10.0, (gain / 40.0))

	return &Filter{
		name: Peaking,
		a0:   1.0 + alpha/a,
		a1:   -2.0 * math.Cos(w0),
		a2:   1.0 - alpha/a,
		b0:   1.0 + alpha*a,
		b1:   -2.0 * math.Cos(w0),
		b2:   1.0 - alpha*a,
	}
}
