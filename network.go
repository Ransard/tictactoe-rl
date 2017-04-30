package main

import (
	"fmt"
	"math"
)

type Network struct {
	Input  [9]Neuron
	Hidden [16]Neuron
	Output [9]Neuron
}

func (n *Network) GetMove(b Board) [9]Neuron {
	fmt.Println("Getting a new move")
	for index, element := range n.Input {
		element.Value = float64(b.Spaces[index])
	}

	n.feedForward()

	return n.Output
}

func (n *Network) feedForward() {
	fmt.Println("Being smart..")

	for ix, eh := range n.Hidden {
		sum := 0.0
		for _, ei := range n.Input {
			sum += ei.Value*eh.Weight + eh.Bias
		}

		n.Hidden[ix].Value = sigmoid(sum)
	}

	for ix, eo := range n.Output {
		sum := 0.0
		for _, eh := range n.Hidden {
			sum += eh.Value*eo.Weight + eo.Bias
		}
		n.Output[ix].Value = sigmoid(sum)
	}
}

func sigmoid(value float64) float64 {
	return value / math.Sqrt(1.0+value*value)
}

func (n *Network) PrintValues() {
	for _, element := range n.Input {
		element.Print()
	}

	for _, element := range n.Hidden {
		element.Print()
	}

	for _, element := range n.Output {
		element.Print()
	}
}

func (n *Network) Init() {
	fmt.Println("Initializing network")

	for ix, _ := range n.Input {
		n.Input[ix].Init(ix)
	}

	for ix, _ := range n.Hidden {
		n.Hidden[ix].Init(ix)
	}

	for ix, _ := range n.Output {
		n.Output[ix].Init(ix)
	}

}
