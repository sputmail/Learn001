package main

func learn(ml *model, times int) int16 {
	for i := 0; i < times; i++ {
		learnmodel(ml)
	}
	return -1
}

func learnmodel(ml *model) {
	learnConnections(ml, ml.Outputs)
	learnlayer(ml, ml.Layer2)
	learnlayer(ml, ml.Layer1)
}

func learnlayer(ml *model, l []node) {
	for i := range l {
		learnConnections(ml, l[i].Connections)
	}
}

func learnConnections(ml *model, cn []connection) {
	for i := range cn {
		learnConnectionByGroup(ml, &cn[i])
	}
}

func learnConnection(ml *model, cn *connection) int {
	wasdeviation := ml.deviation
	wasanswer := ml.answer
	wasvalue := cn.Weight

	cn.Weight = rnd127()
	_, deviationrandom := calculate(ml)
	if deviationrandom < wasdeviation {
		return 0
	} else {
		ml.deviation = wasdeviation
		ml.answer = wasanswer
		cn.Weight = wasvalue

	}

	cn.Weight++
	_, deviationplus := calculate(ml)
	if deviationplus < wasdeviation {

		return 0
	}
	cn.Weight -= 2
	_, deviationminus := calculate(ml)
	if deviationminus < wasdeviation {
		//fmt.Printf(">deviation minus = %d original deviation = %d", deviationminus, wasdeviation)
		return 0
	}

	ml.deviation = wasdeviation
	ml.answer = wasanswer
	cn.Weight = wasvalue
	return 0
}

func learnConnectionByGroup(ml *model, cn *connection) int {
	wasdeviation := ml.GroupSqDeviation
	wasweight := cn.Weight

	cn.Weight = wasweight + 1
	//fmt.Printf("try plus\n")
	deviationplus := calculateAll(ml)
	if deviationplus < wasdeviation {
		//fmt.Printf(">deviation plus = %d original deviation = %d\n", deviationplus, wasdeviation)
		return 0
	}

	cn.Weight = wasweight - 1
	//fmt.Printf("try minus\n")
	deviationminus := calculateAll(ml)
	if deviationminus < wasdeviation {
		//fmt.Printf(">deviation minus = %d original deviation = %d\n", deviationminus, wasdeviation)
		calculateAll(ml)
		return 0
	}

	cn.Weight = rnd127()
	//fmt.Printf("try random\n")
	deviationrandom := calculateAll(ml)
	if deviationrandom < wasdeviation {
		//fmt.Printf(">deviation random = %d original deviation = %d\n", deviationrandom, wasdeviation)
		return 0
	}
	//fmt.Printf(">No change\n")
	cn.Weight = wasweight
	ml.GroupSqDeviation = wasdeviation
	//fmt.Printf(")))))))))))))))))))))) DEFAULT ))))))))))))))))))\n")
	return 0
}
