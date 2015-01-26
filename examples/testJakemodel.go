package main

import (
	"../smallscale"
	"fmt"
	"github.com/istdev/fadingModels/PDPF"
	"github.com/wiless/vlib"
	"math"
	// "matrix"
)

var matlab *vlib.Matlab

func main() {
	nSamples := 10000
	fs := (7.6923e+06) // fs is sampling frequency and fft size = 512
	ts := 1 / fs
	fd := smallscale.GetDoppler(2.0e9, 200)
	fmt.Println(fd)
	pdpm := pdp.PDPManager{}
	pdpm.Load("pdpChannels.json")
	pdp, nTaps := pdpm.GetPDP(ts, "PB")
	
	fmt.Println(pdp)
	pos := vlib.NewVectorI(nTaps)
	j := 0
	for i := 0; i < len(pdp); i++ {
		if pdp[i] > -1.0e3 {
			pdp[i] = math.Pow(10.0, pdp[i]/10.0)
			pos[j] = i
			j++
		} else {
			pdp[i] = 0.0
		}
	}

	matlab = vlib.NewMatlab("test.m")
	channelVal := vlib.NewMatrixC(nSamples, nTaps)
	hn := smallscale.NewMultiTapFading(pdp, fs, fd)
	// hn := ChanFading.MultiTapFading(float64(210))
	for t := 0; t < nSamples; t++ {
		channelVal[t] = hn.Generate(float64(t))
	}
	channelVal1 := channelVal.T()
	fmt.Println(channelVal1.Size())
	// for i := 0; i < nTaps; i++ {
	// 	s[i] = math.Pow(10.0, s[i]/10.0)
	// 	channelVal1[i] = channelVal1[i].Scale(s[i])
	// }

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
	// matlab.Export("t", pos)
	matlab.Command("h = reshape(h, h2, h1);close all;plot(abs(h))")
	matlab.Close()
	fmt.Println("..bye bye")
}
