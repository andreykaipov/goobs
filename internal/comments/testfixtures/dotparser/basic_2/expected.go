package testfixtures

type IDK struct {
	A *A `json:"A,omitempty"`

	C *C `json:"C,omitempty"`
}

type A struct {
	B string `json:"B"`
}

type C struct {
	D string `json:"D"`
}
