package messages

type GameMovePayload struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (m *GameMovePayload) IsValid() bool {
	if m.X >= 3 || m.X < 0 {
		return false
	}

	if m.Y >= 3 || m.Y < 0 {
		return false
	}

	return true
}

type GameMove = Message[GameMovePayload]

func GameMoveFromBytes(bytes []byte) (*GameMove, error) {
	return FromBytes[GameMovePayload]("game.move", bytes)
}
