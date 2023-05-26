package model

import "errors"

var (
	ErrInputNotFound = errors.New("input not found")
)

type Data struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}
