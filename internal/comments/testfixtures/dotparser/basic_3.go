package testfixtures

type Ugh struct {
	// hello
	A string `json:"a"`

	// hi
	B string `json:"b"`

	C struct {
		// bye
		D int `json:"d"`
	} `json:"c"`
}
