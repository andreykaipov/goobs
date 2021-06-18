package testfixtures

type Embedded1 struct {
	A A `json:"a"`
}

type A struct {
	EmbeddedDummy
}
