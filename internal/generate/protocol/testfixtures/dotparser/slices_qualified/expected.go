package testfixtures

import "bytes"

type QualifiedCrap struct {
	ExplicitSlice []bytes.Buffer `json:"explicitSlice"`

	ImplicitSlice []bytes.Buffer `json:"implicitSlice"`
}
