package internal

import "math/rand/v2"

type Room struct {
	playerX       *Client
	playerO       *Client
	board         [3][3]rune
	isPlayerXTurn bool
}

func NewRoom(playerX *Client, playerO *Client) *Room {
	return &Room{
		playerX: playerX,
		playerO: playerO,
		board: [3][3]rune{
			{'_', '_', '_'},
			{'_', '_', '_'},
			{'_', '_', '_'},
		},
		isPlayerXTurn: rand.IntN(10) >= 5,
	}
}

func (r *Room) Run() {
	for {
		select {
		case message, ok := <-r.playerX.Read():
			if !r.isPlayerXTurn {
				r.playerX.Write([]byte("Not your turn"))
				break
			}

			if !ok {
				return
			}

			r.playerO.Write(message)
			r.isPlayerXTurn = !r.isPlayerXTurn
		case message, ok := <-r.playerO.Read():
			if r.isPlayerXTurn {
				r.playerO.Write([]byte("Not your turn"))
				break
			}

			if !ok {
				return
			}

			r.playerX.Write(message)
			r.isPlayerXTurn = !r.isPlayerXTurn
		}
	}
}
