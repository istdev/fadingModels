package main

import (
	"fmt"
	"github.com/wiless/gocomm"
	"github.com/wiless/vlib"

	"../smallscale"
)

////// test multipath with slow fading

func main() {

	var channel smallscale.MPChannel
	channel.InitializeChip()
	N := 100

	param := smallscale.NewIIDChannel()
	fd := smallscale.GetDoppler(speed, fc)
	param := smallscale.NewSlowFadingChannel(fs, fd)

	param.Ts = 4
	pdp := vlib.VectorF{1, .1}
	param.SetPDP(pdp)
	param.Mode = ""
	channel.InitParam(param)
	// samples := vlib.VectorC(sources.RandNCVec(N, 1))
	samples := vlib.NewOnesC(N)

	var data gocomm.SComplex128Obj
	data.Ts = 2
	for i := 0; i < N; i++ {
		data.Ch = samples[i]

		// fmt.Printf("\n Input %d = %v", i, data)
		chout := channel.ChannelFn(data)
		fmt.Printf("\n  %d I/O : %v ==> %v", i, data.Ch, chout.Ch)
		data.UpdateTimeStamp()
	}

}

////////////////
