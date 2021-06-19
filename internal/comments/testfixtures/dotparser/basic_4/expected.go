package testfixtures

import "bytes"

type Basic4 struct {
	C *C `json:"c,omitempty"`

	X []int `json:"x,omitempty"`

	bytes.Buffer
}

type C struct {
	D int `json:"d,omitempty"`
}
