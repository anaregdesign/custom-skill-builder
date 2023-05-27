package skill

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type Book struct {
	m map[string]func(context.Context, []byte) ([]byte, error)
}

func NewBook() *Book {
	return &Book{m: make(map[string]func(context.Context, []byte) ([]byte, error))}
}

func (b *Book) Register(name string, skill func(context.Context, []byte) ([]byte, error)) {
	b.m[name] = skill
}

func (b *Book) Apply(ctx context.Context, name string, body []byte) ([]byte, error) {
	if skill, ok := b.m[name]; ok {
		return skill(ctx, body)
	} else {
		return []byte{}, ErrNotFound
	}
}
