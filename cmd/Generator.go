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
	answer      int8
}

type model struct {
	Version          int
	timeslots        []timeslot
	inputs           []node
	Layer1           []node
	Layer2           []node
	Outputs          []connection
	answer           int16
	rightanswer      int16
	deviation        uint16
	GroupSqDeviation uint32
}

func generateModel() model {
	var m model
	m.Version = 2
	m.inputs = make([]node, countofinputs)
	m.Layer1 = make([]node, countOfNodesLayer1)
	fillLayer(m.Layer1, m.inputs)
	m.Layer2 = make([]node, countOfNodesLayer2)
	fillLayer(m.Layer2, m.Layer1)
	m.Outputs = make([]connection, countOfNodesLayer2)
	fillConnections(m.Outputs, m.Layer2)
	return m
}

func fillLayer(l []node, p []node) {
	for i := range l {
		l[i].Connections = make([]connection, len(p))
		fillConnections(l[i].Connections, p)

	}
}

func fillConnections(c []connection, p []node) {
	for j := range c {
		c[j].parentnode = &p[j]
		c[j].Weight = rnd127() / 5
	}
}

func rnd127() int8 { return int8(rand.Intn(255) - 127) }
