package main

import (
	"math"
)

func testInput(ml *model) int16 {
	var ts timeslot = timeslot{-16, 9, -27, 255, 5, 58, 32, -11}
	inputs := fillinputs(ts)
	answ, _ := calculate(ml, inputs, 0)
	return answ
}

func calculateAll(ml *model) uint32 {
	ml.GroupSqDeviation = 0
	for i := range ml.timeslots {
		inputs := fillinputs(ml.timeslots[i])
		_, dev := calculate(ml, inputs, ml.timeslots[i].answer)
		ml.GroupSqDeviation += uint32(dev) * uint32(dev)
		//fmt.Printf("dev = %d dev2= %d grDeviation = %d\n", dev, uint32(dev)*uint32(dev), ml.GroupSqDeviation)
	}
	//fmt.Printf("________________________________________ grDeviation = %d\n", ml.GroupSqDeviation)
	return ml.GroupSqDeviation
}

func calculate(ml *model, inputs []int8, rightanswer int16) (int16, uint16) {
	answers1 := calcLayer(ml.Layer1, inputs)
	answers2 := calcLayer(ml.Layer2, answers1)
	answer := calcOutput(ml.Outputs, answers2)
	//fmt.Printf("new answer = %d right = %d\n", ml.answer, ml.rightanswer)
	deviation := deviation(answer, rightanswer)
	return answer, deviation
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

func fillinputs(ts timeslot) []int8 {
	inputs := make([]int8, countofinputs)
	inputs[0] = ts.open
	inputs[1] = ts.max
	inputs[2] = ts.min
	inputs[3] = int8(ts.volume - 128)
	inputs[4] = ts.open1
	inputs[5] = ts.open2
	inputs[6] = ts.open3
	return inputs
	//ml.answer, ml.deviation = calculate(ml)
}

func calcLayer(n []node, prev []int8) []int8 {
	answers := make([]int8, len(n))
	for i := range n {
		answers[i] = calcNode(&n[i], prev)
	}
	return answers
}

func calcNode(n *node, prev []int8) int8 {
	var answ int32
	for i := range n.Connections {
		answ += int32(n.Connections[i].Weight) * int32(prev[i])
	}
	//fmt.Printf("sum = %d", answ)
	return int8(sigma(answ/127, 127))
	//fmt.Printf("answer = %d\n", n.answer)
}

func calcOutput(cn []connection, prev []int8) int16 {
	var answ int32
	for i := range cn {
		answ += int32(cn[i].Weight) * int32(prev[i])
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
