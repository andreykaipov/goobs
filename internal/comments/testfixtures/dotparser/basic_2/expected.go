package testfixtures

type IDK struct {
	A A `json:"A"`

	C C `json:"C"`
}

type A struct {
	B string `json:"B"`
}

type C struct {
	D string `json:"D"`
}
