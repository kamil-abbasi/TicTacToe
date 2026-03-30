package messages

import (
	"encoding/json"
	"fmt"
)

type PlayerMoveMessage struct {
	X int
	Y int
}

func PlayerMoveFromBytes(bytes []byte) (*PlayerMoveMessage, error) {
	type Data struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	var data Data

	err := json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal message from bytes")
	}

	return &PlayerMoveMessage{
		X: data.X,
		Y: data.Y,
	}, nil
}

func (m *PlayerMoveMessage) IsValid() bool {
	if m.X >= 3 || m.X < 0 {
		return false
	}

	if m.Y >= 3 || m.Y < 0 {
		return false
	}

	return true
}

func (m *PlayerMoveMessage) ToJsonString() (string, error) {
	data := map[string]int{"x": m.X, "y": m.Y}

	bytes, err := json.Marshal(data)

	if err != nil {
		return "", fmt.Errorf("failed to marshal message to json")
	}

	return string(bytes), nil
}
