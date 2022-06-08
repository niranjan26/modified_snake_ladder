package service

import "meesho/model"

type GameInterface interface {
	CreateGame(int, map[int]int, map[int]int, []int) string
	HoldDice(string, int) bool
	RollDiceAndMove(string, int) bool
}

type SnakesAndLadders struct {
	gameMap map[string]*model.Game
}

func NewSnakesAndLaddersService() *SnakesAndLadders {
	return &SnakesAndLadders{
		gameMap: make(map[string]*model.Game),
	}
}

func (s *SnakesAndLadders) CreateGame(boardSize int, snakes, ladders map[int]int, players []int) string {
	playerLimitPerCell := make([]int, boardSize*boardSize+1)
	for i := 0; i < len(playerLimitPerCell); i++ {
		playerLimitPerCell[i] = 1
	}

	game, id := model.CreateGame(boardSize, snakes, ladders, players, playerLimitPerCell)
	s.gameMap[id] = game

	return id
}

func (s *SnakesAndLadders) HoldDice(gameID string, playerID int) bool {
	game := s.gameMap[gameID]
	if game == nil {
		return false
	}

	return game.HoldDice(playerID)
}

func (s *SnakesAndLadders) RollDiceAndMove(gameID string, playerID int) bool {
	game := s.gameMap[gameID]
	if game == nil {
		return false
	}

	return game.RollDiceAndMove(playerID)
}
