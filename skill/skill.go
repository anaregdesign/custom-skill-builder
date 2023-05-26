package skill

import (
	"encoding/json"
	"errors"
)

var (
	ErrParse = errors.New("parse error")
)

type Skill[S, T any] struct {
	mutation func(S) (T, error)
}

func NewSkill[S, T any](mutation func(S) (T, error)) *Skill[S, T] {
	return &Skill[S, T]{
		mutation: mutation,
	}
}

func NewSkillNoErr[S, T any](mutation func(S) T) *Skill[S, T] {
	return &Skill[S, T]{
		mutation: func(s S) (T, error) {
			return mutation(s), nil
		}}
}
func (s *Skill[S, T]) Apply(body Body[S]) Body[T] {
	result := make([]Record[T], len(body.Values))
	for i, record := range body.Values {
		result[i].RecordID = record.RecordID
		value, err := s.mutation(record.Data)
		if err != nil {
			result[i].Errors = append(result[i].Errors, NewMessage(err.Error()))
		} else {
			result[i].Data = value
		}
	}
	return Body[T]{Values: result}
}

func (s *Skill[S, T]) Flatten() func([]byte) ([]byte, error) {
	return func(body []byte) ([]byte, error) {
		var b Body[S]
		if err := json.Unmarshal(body, &b); err != nil {
			return []byte{}, ErrParse
		}
		result := s.Apply(b)
		if response, err := json.Marshal(result); err != nil {
			return []byte{}, err
		} else {
			return response, nil
		}
	}
}
