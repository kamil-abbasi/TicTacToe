package messages

type Message interface {
	IsValid() bool
	ToJsonString() (string, error)
}
