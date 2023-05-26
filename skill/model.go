package skill

type Message struct {
	Value string `json:"message"`
}

func NewMessage(value string) Message {
	return Message{Value: value}
}

type Record[T any] struct {
	RecordID string       `json:"recordId"`
	Data     map[string]T `json:"data,omitempty"`
	Errors   []Message    `json:"errors,omitempty"`
	Warnings []Message    `json:"warnings,omitempty"`
}

type Body[T any] struct {
	Values []Record[T] `json:"values"`
}
