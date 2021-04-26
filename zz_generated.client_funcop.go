// This file has been automatically generated. Don't edit it.

package goobs

type Option func(*Client)

func NewClient(opts ...Option) *Client {
	o := &Client{}

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func WithHost(x string) Option {
	return func(o *Client) {
		o.Host = x
	}
}

func WithPort(x int) Option {
	return func(o *Client) {
		o.Port = x
	}
}
