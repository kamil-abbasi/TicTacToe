package internal

import (
	"fmt"
	"math/rand/v2"

	"github.com/kamil-abbasi/TicTacToe.git/internal/messages"
)

type Room struct {
	playerX     *Client
	playerO     *Client
	board       [3][3]rune
	currentTurn rune
}

func NewRoom(playerX *Client, playerO *Client) *Room {
	var turn rune
	randomInt := rand.IntN(10)

	if randomInt >= 5 {
		turn = 'x'
	} else {
		turn = 'o'
	}

	return &Room{
		playerX: playerX,
		playerO: playerO,
		board: [3][3]rune{
			{'_', '_', '_'},
			{'_', '_', '_'},
			{'_', '_', '_'},
		},
		currentTurn: turn,
	}
}

func (r *Room) Run() {
	bytes, _ := messages.NewInfo("your turn", "turn").ToBytes()

	if r.currentTurn == 'x' {
		r.playerX.Write(bytes)
	} else {
		r.playerO.Write(bytes)
	}

	for {
		select {
		case message, ok := <-r.playerX.Read():
			if !ok {
				return
			}

			r.processMessage(message, r.playerX, r.playerO, 'x')
		case message, ok := <-r.playerO.Read():
			if !ok {
				return
			}

			r.processMessage(message, r.playerO, r.playerX, 'o')
		}
	}
}

func (r *Room) processMessage(message []byte, player *Client, opponent *Client, playerSymbol rune) {
	if playerSymbol != r.currentTurn {
		bytes, _ := messages.NewInfo("not your turn", "turn").ToBytes()
		player.Write(bytes)
		return
	}

	gameMoveMessage, err := messages.GameMoveFromBytes(message)

	if err != nil {
		bytes, _ := messages.NewError("invalid message format", "validation").ToBytes()
		player.Write(bytes)
		return
	}

	valid := gameMoveMessage.Payload.IsValid()

	if !valid {
		bytes, _ := messages.NewError("invalid message data", "validation").ToBytes()
		player.Write(bytes)
		return
	}

	x := gameMoveMessage.Payload.X
	y := gameMoveMessage.Payload.Y

	if r.board[x][y] != '_' {
		bytes, _ := messages.NewInfo(fmt.Sprintf("square [%v,%v] occupied", x, y), "occupied").ToBytes()
		player.Write(bytes)
		return
	}

	r.board[x][y] = playerSymbol

	bytes, _ := gameMoveMessage.ToBytes()

	opponent.Write(bytes)

	r.checkForWin(player, opponent, playerSymbol)
	r.changeTurn()
}

func (r *Room) changeTurn() {
	if r.currentTurn == 'x' {
		r.currentTurn = 'o'
	} else {
		r.currentTurn = 'x'
	}
}

func (r *Room) checkForWin(player *Client, opponent *Client, playerSymbol rune) {
	gameFinished := false

	// vertical
	for i := range 3 {
		verticalWin := true

		for j := range 3 {
			if r.board[i][j] != playerSymbol {
				verticalWin = false
				break
			}
		}

		if verticalWin {
			gameFinished = true
			break
		}
	}

	// horizontal
	for i := range 3 {
		horizontalWin := true

		for j := range 3 {
			if r.board[j][i] != playerSymbol {
				horizontalWin = false
				break
			}
		}

		if horizontalWin {
			gameFinished = true
			break
		}
	}

	// first diagonal
	diagonalWin := true

	for i := range 3 {
		if r.board[i][i] != playerSymbol {
			diagonalWin = false
			break
		}
	}

	for i := 2; i <= 0; i-- {
		if r.board[i][i] != playerSymbol {
			diagonalWin = false
			break
		}
	}

	if diagonalWin {
		gameFinished = true
	}

	if gameFinished {
		bytes1, _ := messages.NewInfo("you won", "victory").ToBytes()
		bytes2, _ := messages.NewInfo("you lost", "loss").ToBytes()
		player.Write(bytes1)
		opponent.Write(bytes2)

		player.Close()
		opponent.Close()
	}
}
