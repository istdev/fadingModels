package main

import (
	"github.com/istdev/fadingModels/PDP"
	"github.com/istdev/fadingModels/smallscalechan"

	//"../smallscale"
	"fmt"
	"github.com/wiless/gocomm"
	"github.com/wiless/vlib"
)

////// test multipath with slow fading

func main() {

	var channel smallscalechan.MPChannel
	channel.InitializeChip()
	N := 5
	fs := 7.6e6 // In Hz
	fc := 1.8e3 // In MHz

	fd := smallscalechan.GetDoppler(fc, 5)
	fmt.Println(fd)
	// param := smallscalechan.NewSlowFadingChannel(fs, fd)
	param := smallscalechan.NewIIDChannel()

	param.Ts = -1
	// pdp := vlib.VectorF{1, .1}
	// param.SetPDP(pdp)
	// param.Mode = ""

	pdpm := pdp.PDPManager{}
	pdpm.Load("pdpChannels.json")
	pdp, _ := pdpm.GetPDPLinear(1/fs, "PB")
	param.SetPDP(pdp)
	fmt.Println("The PDP is", pdp, "with Other params as ", param.Ts, param.Mode, param.FdopplerHz)
	channel.InitParam(param)
	// samples := vlib.VectorC(sources.RandNCVec(N, 1))
	samples := vlib.NewOnesC(N)

	var data gocomm.SComplex128Obj
	data.Ts = 1 / fs
	for i := 0; i < N; i++ {
		data.Ch = samples[i]

		// fmt.Printf("\n Input %d = %v", i, data)
		chout := channel.ChannelFn(data)
		fmt.Printf("\n  %d I/O : %v ==> %v", i, data.Ch, chout.Ch)
		data.UpdateTimeStamp()
	}

}

////////////////
