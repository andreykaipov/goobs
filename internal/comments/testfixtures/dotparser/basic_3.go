package testfixtures

type Ugh struct {
	// hello
	A string `json:"a"`

	// hi
	B string `json:"b"`

	C C `json:"c"`
}

type C struct {
	// bye
	D int `json:"d"`
}
