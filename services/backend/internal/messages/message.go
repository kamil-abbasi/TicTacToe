package messages

import (
	"encoding/json"
	"fmt"
)

type MessageType struct {
	str string
}

func NewType(messageType string) *MessageType {
	return &MessageType{
		str: messageType,
	}
}

func (t *MessageType) AddSubtype(subtype string) *MessageType {
	if subtype == "" {
		return t
	}

	t.str = fmt.Sprintf("%v.%v", t.str, subtype)

	return t
}

func (t *MessageType) String() string {
	return t.str
}

type Message[Payload any] struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

func New[Payload any](messageType string, payload Payload) *Message[Payload] {
	return &Message[Payload]{
		Type:    messageType,
		Payload: payload,
	}
}

func FromBytes[Payload any](messageType string, bytes []byte) (*Message[Payload], error) {
	var payload Payload

	err := json.Unmarshal(bytes, &payload)

	if err != nil {
		return &Message[Payload]{}, err
	}

	return New(messageType, payload), nil
}

func (m *Message[Payload]) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(m)

	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
