package skill

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Book struct {
	m map[string]func([]byte) ([]byte, error)
}

func NewBook() *Book {
	return &Book{m: make(map[string]func([]byte) ([]byte, error))}
}

func (b *Book) Register(name string, skill func([]byte) ([]byte, error)) {
	b.m[name] = skill
}

func (b *Book) Apply(name string, body []byte) ([]byte, error) {
	if skill, ok := b.m[name]; ok {
		return skill(body)
	} else {
		return []byte{}, ErrNotFound
	}
}
