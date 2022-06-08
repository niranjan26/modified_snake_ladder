package model

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"

	"github.com/google/uuid"
)

type Game struct {
	id       string
	board    *Board
	players  map[int]*Player
	winner   *Player
	isOver   bool
	diceLock sync.Mutex
	diceHold *Player
}

type Board struct {
	boardSize int
	cells     map[int]*Cell
	lastCell  *Cell
	startCell *Cell
}

type Cell struct {
	id          int
	players     map[int]*Player
	playerLimit int
	isSnake     bool
	snakeDst    *Cell
	isLadder    bool
	ladderDst   *Cell
}

type Player struct {
	id   int
	cell int
}

func CreateGame(boardSize int, snakes, ladders map[int]int, players []int, playerLimitPerCell []int) (*Game, string) {
	game := &Game{
		id:      uuid.NewString(),
		board:   CreateBoard(boardSize, len(players), playerLimitPerCell),
		players: CreatePlayers(players),
		isOver:  false,
	}

	playersMap := game.players
	game.board.cells[0].players = playersMap

	updateSnakesAndLadders(game.board, snakes, ladders)

	return game, game.id
}

func (g *Game) isPlayerPartOfGame(player int) bool {
	if _, ok := g.players[player]; !ok {
		return false
	}

	return true
}

func (g *Game) valid(player int) bool {
	// g.Printer(g.players[player])
	if !g.isPlayerPartOfGame(player) {
		return false
	}

	if g.isOver {
		return false
	}

	return true
}

func (g *Game) Printer(player *Player) {
	fmt.Printf("game: %s, player: %v, isOver: %v\n", g.id, *player, g.isOver)
	if g.diceHold != nil {
		fmt.Printf("game: %s, player: %v, diceHold: %v\n", g.id, *player, *g.diceHold)
	}
	for _, cell := range g.board.cells {
		fmt.Printf("game: %s, cell:%v, playerInThisCell: %v\n", g.id, *cell, cell.players)
	}
}

func (g *Game) HoldDice(player int) bool {
	if !g.valid(player) {
		return false
	}

	if g.diceHold != nil {
		fmt.Println(*g.diceHold)
		return false
	}

	g.diceLock.Lock()
	if g.diceHold != nil {
		g.diceLock.Unlock()
		return false
	}

	g.diceHold = g.players[player]
	g.diceLock.Unlock()

	return true
}

func (g *Game) RollDiceAndMove(player int) bool {
	if !g.valid(player) {
		return false
	}

	if g.diceHold != g.players[player] {
		return false
	}

	roll, _ := rand.Int(rand.Reader, big.NewInt(6))
	num := roll.Int64()

	if !g.makeMove(int(num), player) {
		return false
	}

	return true
}

func (g *Game) makeMove(num, playerID int) bool {
	player := g.players[playerID]
	cellID := player.cell
	cell := g.board.cells[cellID]

	if cellID+num > g.board.boardSize {
		return false
	}

	tempCell := g.board.cells[cellID+num]

	if tempCell.isSnake {
		tempCell = tempCell.snakeDst
	}

	if tempCell.isLadder {
		tempCell = tempCell.ladderDst
	}

	if len(tempCell.players) >= tempCell.playerLimit {
		return false
	}

	if tempCell.id == g.board.boardSize {
		g.isOver = true
		g.winner = player
	}

	delete(cell.players, playerID)
	tempCell.players[playerID] = player

	g.diceHold = nil

	return true
}

func updateSnakesAndLadders(board *Board, snakes, ladders map[int]int) {
	for key, val := range snakes {
		board.cells[key].isSnake = true
		board.cells[key].snakeDst = board.cells[val]
	}

	for key, val := range ladders {
		board.cells[key].isLadder = true
		board.cells[key].ladderDst = board.cells[val]
	}
}

func CreateBoard(boardSize, playersCount int, playerLimit []int) *Board {
	cells := make(map[int]*Cell)
	for i := 0; i < boardSize*boardSize+1; i++ {
		cells[i] = &Cell{
			id:          i,
			players:     make(map[int]*Player),
			playerLimit: playerLimit[i],
			isSnake:     false,
			isLadder:    false,
		}
	}

	cells[0].playerLimit = playersCount

	return &Board{
		boardSize: boardSize * boardSize,
		cells:     cells,
		lastCell:  cells[boardSize*boardSize],
		startCell: cells[0],
	}
}

func CreatePlayers(playersList []int) map[int]*Player {
	players := make(map[int]*Player)
	for _, id := range playersList {
		players[id] = &Player{id: id, cell: 0}
	}

	return players
}
