package model

import "errors"

var (
	ErrInputNotFound = errors.New("input not found")
)

type StringData struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}

type IntData struct {
	Input  int `json:"input,omitempty"`
	Output int `json:"output,omitempty"`
}

type FloatData struct {
	Input  float64 `json:"input,omitempty"`
	Output float64 `json:"output,omitempty"`
}

type BoolData struct {
	Input  bool `json:"input,omitempty"`
	Output bool `json:"output,omitempty"`
}

type CollectionStringData struct {
	Input  []string `json:"input,omitempty"`
	Output []string `json:"output,omitempty"`
}

type CollectionIntData struct {
	Input  []int `json:"input,omitempty"`
	Output []int `json:"output,omitempty"`
}

type CollectionFloatData struct {
	Input  []float64 `json:"input,omitempty"`
	Output []float64 `json:"output,omitempty"`
}

type CollectionBoolData struct {
	Input  []bool `json:"input,omitempty"`
	Output []bool `json:"output,omitempty"`
}
