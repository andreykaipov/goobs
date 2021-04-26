package testfixtures

type Embedded1 struct {
	A struct {
		EmbeddedDummy
	} `json:"a"`
}
