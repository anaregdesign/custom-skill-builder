package skill

type message struct {
	Value string `json:"message"`
}

func newMessage(value string) message {
	return message{Value: value}
}

type record[T any] struct {
	RecordID string    `json:"recordId"`
	Data     T         `json:"data,omitempty"`
	Errors   []message `json:"errors,omitempty"`
	Warnings []message `json:"warnings,omitempty"`
}

type body[T any] struct {
	Values []record[T] `json:"values"`
}
