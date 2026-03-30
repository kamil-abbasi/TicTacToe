package messages

import (
	"encoding/json"
	"fmt"
)

type ErrorMessage struct {
	Message string
}

func NewError(message string) Message {
	return &ErrorMessage{
		Message: message,
	}
}

func (m *ErrorMessage) IsValid() bool {
	return true
}

func (m *ErrorMessage) ToJsonString() (string, error) {
	data := map[string]string{"message": m.Message}

	bytes, err := json.Marshal(data)

	if err != nil {
		return "", fmt.Errorf("failed to marshal message to json")
	}

	return string(bytes), nil
}
