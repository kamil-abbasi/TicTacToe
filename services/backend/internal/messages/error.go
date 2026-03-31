package messages

type ErrorPayload struct {
	Error string `json:"error"`
}

type Error = Message[ErrorPayload]

func NewError(errorMessage string, subtype string) *Error {
	messageType := NewType("error").AddSubtype(subtype).String()

	return New(messageType, ErrorPayload{
		Error: errorMessage,
	})
}
