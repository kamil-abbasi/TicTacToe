package messages

type InfoPayload struct {
	Message string `json:"message"`
}

type Info = Message[InfoPayload]

func NewInfo(message string, subtype string) *Info {
	messageType := NewType("info").AddSubtype(subtype).String()

	return New(messageType, InfoPayload{
		Message: message,
	})
}
