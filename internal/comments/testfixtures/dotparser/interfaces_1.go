package testfixtures

type Interfaces1 struct {
	A A `json:"a"`
}

type A struct {
	MyInterface Interface
}
