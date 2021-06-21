package testfixtures

type Interfaces1 struct {
	A *A `json:"a,omitempty"`
}

type A struct {
	MyInterface Interface
}
