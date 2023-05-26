package skill

import (
	"encoding/json"
	"errors"
)

var (
	ErrParse              = errors.New("parse error")
	ErrInputFieldNotFound = errors.New("input field not found")
)

type Skill[S, T any] struct {
	inputFieldName  string
	outputFieldName string
	mutation        func(S) (T, error)
}

func NewSkill[S, T any](input string, output string, mutation func(S) (T, error)) *Skill[S, T] {
	return &Skill[S, T]{
		inputFieldName:  input,
		outputFieldName: output,
		mutation:        mutation,
	}
}

func NewSkillNoErr[S, T any](input string, output string, mutation func(S) T) *Skill[S, T] {
	return &Skill[S, T]{
		inputFieldName:  input,
		outputFieldName: output,
		mutation: func(s S) (T, error) {
			return mutation(s), nil
		}}
}

func (s *Skill[S, T]) Apply(body Body[S]) Body[T] {
	result := make([]Record[T], len(body.Values))
	for i, record := range body.Values {
		result[i].RecordID = record.RecordID
		result[i].Data = make(map[string]T)
		v, ok := record.Data[s.inputFieldName]
		if !ok {
			result[i].Errors = append(result[i].Errors, NewMessage(ErrInputFieldNotFound.Error()))
			continue
		}
		value, err := s.mutation(v)
		if err != nil {
			result[i].Errors = append(result[i].Errors, NewMessage(err.Error()))
		} else {
			result[i].Data[s.outputFieldName] = value
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
