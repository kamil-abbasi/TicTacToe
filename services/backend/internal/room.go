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
	if r.currentTurn == 'x' {
		r.playerX.Write([]byte("Your turn"))
	} else {
		r.playerO.Write([]byte("Your turn"))
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
		player.Write([]byte("Not your turn"))
		return
	}

	playerMoveMessage, err := messages.PlayerMoveFromBytes(message)

	if err != nil {
		player.Write([]byte("invalid message format"))
		return
	}

	valid := playerMoveMessage.IsValid()

	if !valid {
		player.Write([]byte("invalid message data"))
		return
	}

	jsonString, err := playerMoveMessage.ToJsonString()

	if err != nil {
		player.Write([]byte("invalid message format"))
		return
	}

	if r.board[playerMoveMessage.X][playerMoveMessage.Y] != '_' {
		player.Write([]byte(
			fmt.Sprintf("square [%v,%v] occupied", playerMoveMessage.X, playerMoveMessage.Y),
		))

		return
	}

	r.board[playerMoveMessage.X][playerMoveMessage.Y] = playerSymbol

	opponent.Write([]byte(jsonString))

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
		player.Write([]byte("You won"))
		opponent.Write([]byte("You lost"))

		player.Close()
		opponent.Close()
	}
}
