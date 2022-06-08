package main

import (
	"fmt"
	"meesho/service"
	"sync"
)

func main() {

	snl := service.NewSnakesAndLaddersService()

	boardSize := 5
	ladders := map[int]int{
		1: 21,
		2: 22,
		3: 23,
		4: 24,
		5: 25,
	}

	snakes := map[int]int{}

	game := snl.CreateGame(boardSize, ladders, snakes, []int{1, 2, 3})

	fmt.Printf("1 holddice:%v\n", snl.HoldDice(game, 1))
	fmt.Printf("2 holddice:%v\n", snl.HoldDice(game, 2))
	fmt.Printf("1 rolDice:%v\n", snl.RollDiceAndMove(game, 1))
	fmt.Printf("1 rolDice:%v\n", snl.RollDiceAndMove(game, 1))
	fmt.Printf("2 rolDice:%v\n", snl.RollDiceAndMove(game, 2))
	fmt.Printf("2 holddice:%v\n", snl.HoldDice(game, 2))
	fmt.Printf("2 rolDice:%v\n", snl.RollDiceAndMove(game, 2))

	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go check(&wg, i, snl, game)
	// }

	// wg.Wait()
}

func check(wg *sync.WaitGroup, i int, snl *service.SnakesAndLadders, game string) {
	defer wg.Done()

	players := []int{1, 2, 3}
	for _, player := range players {
		fmt.Printf("thread: %d, holdDice: %v, rollDiceAndMove: %v\n", i, snl.HoldDice(game, player), snl.RollDiceAndMove(game, player))
	}
}
