package main

import "fmt"

type Game struct {
	Status string
}

func (g Game) Play() {
	ai := Network{}
	currBoard := Board{}
	fmt.Println("Are you ready?!")
	ai.Init()
	currBoard.AI = ai
	result := []int{0, 0, 0}
	for i := 0; i < 500; i++ {
		winner := currBoard.InitSpaces()
		result[winner]++
	}
	fmt.Println("Winner is ", Space(0).ToString(), Space(1).ToString(), Space(2).ToString())
	fmt.Println("Winner is ", result)
	return
}
