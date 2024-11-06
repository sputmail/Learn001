package main

import (
	"math/rand"
)

type connection struct {
	Weight     int8
	parentnode *node
}

type node struct {
	Connections []connection
	//	answer      int8
}

type model struct {
	Version   int
	timeslots []timeslot
	//inputs    []node
	Layer1  []node
	Layer2  []node
	Outputs []connection
	//answer           int16
	//rightanswer      int16
	//deviation        uint16
	GroupSqDeviation uint32
}

func generateMinModel(ts []timeslot) model {
	mindev := uint32(10000)
	m := generateModel()
	filltimeslots(&m, ts, 0, groupSize*100)
	mindev = calculateAll(&m)

	for i := 0; i < 3000; i++ {
		m2 := generateModel()
		filltimeslots(&m2, ts, 0, 3480)
		dev2 := calculateAll(&m2)
		if dev2 < mindev {
			m = m2
			mindev = dev2
		}
	}
	SaveModel(&m, "i")
	return m
}

func generateModel() model {
	var m model
	m.Version = variant
	//m.inputs = make([]node, countofinputs)
	m.Layer1 = make([]node, countOfNodesLayer1)
	fillLayer(m.Layer1, countofinputs)
	m.Layer2 = make([]node, countOfNodesLayer2)
	fillLayer(m.Layer2, countOfNodesLayer1)
	m.Outputs = make([]connection, countOfNodesLayer2)
	fillConnections(m.Outputs)
	return m
}

func fillLayer(l []node, lenPrevLayer int) {
	for i := range l {
		l[i].Connections = make([]connection, lenPrevLayer)
		fillConnections(l[i].Connections)

	}
}

func fillConnections(c []connection) {
	for j := range c {
		//c[j].parentnode = &p[j]
		c[j].Weight = rnd127() / 5
	}
}

func rnd127() int8 { return int8(rand.Intn(255) - 127) }
