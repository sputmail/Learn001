package main

import (
	"math"
)

func testInput(ml *model) int16 {
	var ts timeslot = timeslot{-16, 9, -27, 255, 5, 58, 32, -11}
	fillinputs(ml, ts)
	answ, _ := calculate(ml)
	return answ
}

func calculateAll(ml *model) uint32 {
	ml.GroupSqDeviation = 0
	for i := range ml.timeslots {
		fillinputs(ml, ml.timeslots[i])
		_, dev := calculate(ml)
		ml.GroupSqDeviation += uint32(dev) * uint32(dev)
		//fmt.Printf("dev = %d dev2= %d grDeviation = %d\n", dev, uint32(dev)*uint32(dev), ml.GroupSqDeviation)
	}
	//fmt.Printf("________________________________________ grDeviation = %d\n", ml.GroupSqDeviation)
	return ml.GroupSqDeviation
}

func calculate(ml *model) (int16, uint16) {
	calcLayer(ml.Layer1)
	calcLayer(ml.Layer2)
	ml.answer = calcOutput(ml.Outputs)
	//fmt.Printf("new answer = %d right = %d\n", ml.answer, ml.rightanswer)
	ml.deviation = deviation(ml.answer, ml.rightanswer)
	return ml.answer, ml.deviation
}

func filltimeslots(ml *model, ts []timeslot, begin int, count int) {
	ml.timeslots = nil
	for i := begin; i < begin+count; i++ {
		ind := i
		if i >= len(ts) {
			ind = i - len(ts)
		}
		ml.timeslots = append(ml.timeslots, ts[ind])
	}
	ml.GroupSqDeviation = calculateAll(ml)
}

func fillinputs(ml *model, ts timeslot) {
	ml.inputs[0].answer = ts.open
	ml.inputs[1].answer = ts.max
	ml.inputs[2].answer = ts.min
	ml.inputs[3].answer = int8(ts.volume - 128)
	ml.inputs[4].answer = ts.open1
	ml.inputs[5].answer = ts.open2
	ml.inputs[6].answer = ts.open3
	ml.rightanswer = ts.answer
	ml.answer, ml.deviation = calculate(ml)
}

func calcLayer(n []node) {
	for i := range n {
		//fmt.Printf("calc node %d", i)
		calcNode(&n[i])

	}
}

func calcNode(n *node) {
	var answ int32
	for i := range n.Connections {
		answ += int32(n.Connections[i].Weight) * int32(n.Connections[i].parentnode.answer)
	}
	//fmt.Printf("sum = %d", answ)
	n.answer = int8(sigma(answ/127, 127))
	//fmt.Printf("answer = %d\n", n.answer)
}

func calcOutput(cn []connection) int16 {
	var answ int32
	for i := range cn {
		answ += int32(cn[i].Weight) * int32(cn[i].parentnode.answer)
		//fmt.Printf(" %d: w = %d * a %d  ", i, cn[i].Weight, cn[i].parentnode.answer)
	}
	//fmt.Printf("calc output answ = %d\n", answ)
	return sigma(answ/4, 1024)
}

func sigma(value int32, mlt int32) int16 {
	vl := float64(value) / float64(mlt)
	//fmt.Printf(" value= %d\n", value)
	return int16(vl / math.Sqrt(1+vl*vl) * float64(mlt))
}

func deviation(value, answer int16) uint16 {
	if value > answer {
		return uint16(value - answer)
	} else {
		return uint16(answer - value)
	}
}
