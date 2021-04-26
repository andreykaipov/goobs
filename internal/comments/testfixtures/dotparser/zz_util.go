package testfixtures

// these types exist so this Go package doesn't error out.
// they're used in our tests.

type Interface interface {
	Wow() string
}

type EmbeddedDummy struct {
	a int
	b string
}
