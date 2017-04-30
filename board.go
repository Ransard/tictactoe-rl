package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Board struct {
	Spaces [9]Space
	AI     Network
}

func (b Board) GetSpace(i int) string {
	return b.Spaces[i].ToString()
}

func (b Board) InitSpaces() Space {

	crossTurn := false
	round := 1

	for b.placeNewMarker(crossTurn) {
		//b.printStatus(round)

		gameOver, winner := b.gameIsOver()

		if gameOver {
			return winner
		}

		crossTurn = !crossTurn
		round++
	}

	return FREE
}

func (b *Board) placeNewMarker(crossTurn bool) bool {

	canPlace := false
	for i := 0; i < 9; i++ {
		if b.Spaces[i] == FREE {
			canPlace = true
		}
	}

	if !canPlace {
		return false
	}

	if crossTurn && b.placeRandomMarker() {
		return true
	} else if !crossTurn && b.placeSmartMarker() {
		return true
	}

	return b.placeNewMarker(crossTurn)
}

func (b *Board) placeRandomMarker() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tryIndex := r.Intn(9)

	if b.Spaces[tryIndex] == FREE {
		b.Spaces[tryIndex] = CROSS
		return true
	}

	return false
}

type ByValue [9]Neuron

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

func (b *Board) placeSmartMarker() bool {

	result := b.AI.GetMove(*b)

	sort.Sort(ByValue(result))

	for ix, _ := range result {

		tryIndex := result[ix].ID

		if b.Spaces[tryIndex] == FREE {
			b.Spaces[tryIndex] = NOUGHT
			return true
		}

	}

	return false
}

func (b *Board) gameIsOver() (bool, Space) {

	finishedGame, winner := b.winCondition()

	if finishedGame {
		return finishedGame, winner
	}

	for i := 0; i < 9; i++ {
		if b.Spaces[i] == FREE {
			return false, FREE
		}
	}

	return false, FREE
}

func (b Board) printStatus(round int) {
	fmt.Println("Round ", round)
	fmt.Println("-------------------")
	fmt.Println(b.GetSpace(0), b.GetSpace(1), b.GetSpace(2))
	fmt.Println(b.GetSpace(3), b.GetSpace(4), b.GetSpace(5))
	fmt.Println(b.GetSpace(6), b.GetSpace(7), b.GetSpace(8))
	fmt.Println("")
}

func (b Board) winCondition() (result bool, space Space) {

	result, space = b.checkIfSame(0, 1, 2)

	if !result {
		result, space = b.checkIfSame(3, 4, 5)
	}

	if !result {
		result, space = b.checkIfSame(6, 7, 8)
	}

	if !result {
		result, space = b.checkIfSame(0, 3, 6)
	}

	if !result {
		result, space = b.checkIfSame(1, 4, 7)
	}

	if !result {
		result, space = b.checkIfSame(2, 5, 8)
	}

	if !result {
		result, space = b.checkIfSame(0, 4, 8)
	}

	if !result {
		result, space = b.checkIfSame(2, 4, 6)
	}

	return
}

func (b Board) checkIfSame(x, y, z int) (bool, Space) {

	if b.Spaces[x] == FREE {
		return false, FREE
	}

	result := b.Spaces[x] == b.Spaces[y] && b.Spaces[y] == b.Spaces[z]
	if result {
		fmt.Println("We have a winner! - ", b.GetSpace(x), " - ", x, y, z)
	}

	return result, b.Spaces[x]
}
