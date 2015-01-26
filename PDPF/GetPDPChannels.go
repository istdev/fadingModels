package pdp

import (
	"encoding/json"
	// "fmt"
	"github.com/wiless/vlib"
	"io/ioutil"
	"log"
	"math"
)

type PDP struct {
	Name      string
	TimeINns  vlib.VectorF
	PowerINdb vlib.VectorF
}

type PDPManager struct {
	pdpData []PDP
}

func (p *PDPManager) Load(filename string) bool {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("File Error is ", err)
		return false
	}

	// var u PDP
	jerr := json.Unmarshal(data, &p.pdpData)
	if jerr != nil {
		return false
	}
	return true
}

func (p *PDPManager) Count() int {
	return len(p.pdpData)
}

//GetPDP returns the PDP in Db at sample rate Ts

func (p *PDPManager) GetPDP(Ts float64, pdpName string) (vlib.VectorF, int) {

	// if jerr != nil {
	//     log.Println("Json Unmarshal error is ", jerr)
	// }
	// for i := 0; i < p.Count(); i++ {
	// 	log.Println("Available PDPS", p.pdpData[i].Name)
	// }
	var pp PDP
	for j := 0; j < p.Count(); j++ {
		if p.pdpData[j].Name == pdpName {
			pp = p.pdpData[j]
		}
		// fmt.Printf("\n%d : PDP Info is %v", i, v[i])
	}
	// if pp.Name==""

	// log.Println("Found this PDP  %v", pp)
	Ts = Ts * 1.0e9
	t := vlib.NewVectorF(len(pp.TimeINns))
	s := vlib.NewVectorF(len(pp.PowerINdb))
	t = pp.TimeINns.Scale(1 / Ts)
	nTaps := len(pp.TimeINns)
	timeIndex := vlib.NewVectorI(nTaps)
	for i := 0; i < nTaps; i++ {
		s[i] = pp.PowerINdb[i]
		s[i] = math.Pow(10.0, (s[i] / 10.0))
		// fmt.Println(s[i])
		timeIndex[i] = int(t[i])
	}
	result := (vlib.NewOnesF(timeIndex[nTaps-1] + 1)).Scale(-1.0e5)
	s = s.Scale(1 / (vlib.Sum(s)))
	// fmt.Println(result)
	for i := 0; i < nTaps; i++ {
		s[i] = 10 * math.Log10(s[i])
		result[timeIndex[i]] = pp.PowerINdb[i]
	}
	// t = pp.TimeINns.Scale(1 / Ts)

	return result, nTaps
}

func (p *PDPManager) GetPDPLinear(Ts float64, name string) (pdp vlib.VectorF, nzPos vlib.VectorI) {

	pdp, nTaps := p.GetPDP(Ts, name)

	// fmt.Println(pdp)
	pos := vlib.NewVectorI(nTaps)
	j := 0
	for i := 0; i < len(pdp); i++ {
		if pdp[i] > -1.0e3 {

			pdp[i] = vlib.InvDb(pdp[i]) //math.Pow(10.0, pdp[i]/10.0)
			pos[j] = i
			j++
		} else {
			pdp[i] = 0.0
		}
	}
	return pdp, pos
}

// func main() {

// 	data, err := ioutil.ReadFile("pdpChannels.json")
// 	if err != nil {
// 		log.Println("File Error is ", err)
// 	}

// 	var v []PDP
// 	var u PDP
// 	jerr := json.Unmarshal(data, &v)
// 	if jerr != nil {
// 		log.Println("Json Unmarshal error is ", jerr)
// 	}
// 	for i := 0; i < len(v); i++ {
// 		if v[i].Name == "IndA" {
// 			u = v[i]
// 		}
// 		// fmt.Printf("\n%d : PDP Info is %v", i, v[i])
// 	}
// 	fmt.Printf("\n%d : PDP Info is %v", 0, u)

// }
