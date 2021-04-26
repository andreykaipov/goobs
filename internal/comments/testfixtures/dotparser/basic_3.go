package testfixtures

type Ugh struct {
	A string `json:"a"`
	B string `json:"b"`
	C struct {
		D int `json:"d"`
	} `json:"c"`
}
