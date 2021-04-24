# funcopgen

funcopgen generates [functional
options](https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md)
for your Go structs.

## usage

Better illustrated through example, here's an `Animal` struct with the necessary
`go:generate` directive.

```go
package animal

//go:generate go run github.com/andreykaipov/funcopgen -type=Animal -factory

type Animal struct {
	Surname string `default:"n/a"`
	Color   string `default:"red"`
	cute    bool
}
```

Assuming we've already initialized a Go module, generate!

```console
GOFLAGS=-mod=mod go generate ./...
```

Enjoy the new file `zz_generated.animal_funcop.go` in your package:

```go
// This file has been automatically generated. Don't edit it.

package animal

type Option func(*Animal)

func NewAnimal(opts ...Option) *Animal {
	o := &Animal{
		Color:   "red",
		Surname: "n/a",
	}

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func Color(x string) Option {
	return func(o *Animal) {
		o.Color = x
	}
}

func Surname(x string) Option {
	return func(o *Animal) {
		o.Surname = x
	}
}
```

Now we can instantiate our animals as follows:

```go
db := NewAnimal(
	Surname("ducky"),
	Color("blue"),
)
```

Fun!

### extras

The generated code can be tweaked by passing extra flags:

```console
Usage of funcopgen:
  -factory
        If present, add a factory function for your type, e.g. NewAnimal(opt ...Option)
  -prefix string
        Prefix to attach to functional options, e.g. WithColor, WithName, etc.
  -type string
        Comma-delimited list of type names
  -unexported
        If present, functional options are also generated for unexported fields.
  -unique-option
        If present, prepends the type to the Option type, e.g. AnimalOption.
        Handy if generating for several structs within the same package.
```

See [examples/test.go](./examples/test.go) for an example of all of these flags.

## faq

### I vendor my dependencies. How can I vendor this tool?

In the usage example above, we used `-mod=mod` to tell Go to ignore the vendor
directory since this tool isn't ever imported as a package in any module. If
instead we'd like to vendor this tool, the recommended approach can be found in
detail under [_How can I track tool dependencies for
a module?_](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)
from Go's GitHub wiki.

The TLDR is to add a `tools.go` in your project with the following contents:

```go
// +build tools

package tools

import (
	_ "github.com/andreykaipov/funcopgen"
)
```

After a `go mod vendor`, we can omit the `-mod=mod` we added originally.
However, if your Go version is 1.13 or lower, you'll need to explicitly tell Go
to use the vendored version, i.e. `GOFLAGS=-mod=vendor go generate ./...`.

### How do I integrate it into my development lifecycle?

Code generation shouldn't happen often, but it's easy enough to integrate this
into our build. Just `go generate` before a `go build`. For example, if we're
using make as our build tool:

```Makefile
generate:
    go generate ./...

build: generate
    go build ./...
```
