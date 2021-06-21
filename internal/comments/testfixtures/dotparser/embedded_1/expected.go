package testfixtures

type Embedded1 struct {
	A *A `json:"a,omitempty"`
}

type A struct {
	EmbeddedDummy
}
