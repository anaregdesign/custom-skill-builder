package model

import "errors"

var (
	ErrInputNotFound = errors.New("input not found")
)

type StringData struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}
