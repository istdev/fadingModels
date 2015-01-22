package main

import (
	"../ChanFading"
	"fmt"
	"github.com/wiless/vlib"
	// "math"
	// "matrix"
)

var matlab *vlib.Matlab

func main() {
	nTaps := 3
	nSamples := 200
	fs := 7.6923e+06 / 512 // fs is sampling frequency and fft size = 512
	fd := 50.0
	matlab = vlib.NewMatlab("test.m")
	channelVal := vlib.NewMatrixC(nSamples, nTaps)
	hn := ChanFading.NewMultiTapFading(nTaps, fs, fd)
	// hn := ChanFading.MultiTapFading(float64(210))
	for t := 0; t < nSamples; t++ {
		channelVal[t] = hn.Generate(float64(t))
	}

	channelVal1 := channelVal.T()
	// fmt.Println(channelVal1.Size())
	// channelVal2 := vlib.NewVectorC(nSamples)
	// rms := vlib.NewVectorF(nTaps)
	// for i := 0; i < nTaps; i++ {
	// 	channelVal2 = channelVal1.GetRow(i)
	// 	rms[i] = vlib.Norm2C(channelVal2) / math.Sqrt(float64(nSamples))
	// 	channelVal2 = channelVal2.Scale(1 / rms[i])
	// 	channelVal1[i] = channelVal2
	// }
	// // rms := vlib.Norm2C(channelVal)
	matlab.Export("h", channelVal1)
	matlab.Export("h1", channelVal1.NRows())
	matlab.Export("h2", channelVal1.NCols())
	matlab.Command("h = reshape(h, h2, h1);close all;plot(abs(h)); xlabel('OFDM Symbols'); ylabel('Channel Taps Gain')")
	matlab.Close()
	fmt.Println("..bye bye")
}
