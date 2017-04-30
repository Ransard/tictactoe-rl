package main

import (
	"fmt"
	"math/rand"
)

type Neuron struct {
	ID     int
	Value  float64
	Weight float64
	Bias   float64
}

func (n *Neuron) Init(id int) {
	n.Weight = rand.Float64()
	n.Bias = rand.Float64()
	n.ID = id
}

func (n Neuron) Print() {
	fmt.Println("Weight is ", n.Weight, " Bias is ", n.Bias)
}
