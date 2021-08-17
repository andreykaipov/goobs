package e2e

import (
	"fmt"
	"testing"

	"github.com/andreykaipov/goobs"
	"github.com/stretchr/testify/assert"
)

type customLogger struct{}

func (t customLogger) Printf(format string, v ...interface{}) {}

func Test_goobs_client_options(t *testing.T) {
	var client *goobs.Client
	var err error

	t.Run("basic", func(t *testing.T) {
		client, err = goobs.New("localhost:4545", goobs.WithPassword("hello"))
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%T", client.Log), "*goobs.discard")
	})

	t.Run("default logger used when only debug", func(t *testing.T) {
		client, err = goobs.New("localhost:4545", goobs.WithPassword("hello"), goobs.WithDebug(true))
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%T", client.Log), "*log.Logger")
	})

	t.Run("custom logger implies debug logging", func(t *testing.T) {
		client, err = goobs.New("localhost:4545", goobs.WithPassword("hello"), goobs.WithLogger(customLogger{}))
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%T", client.Log), "e2e.customLogger")
	})

	t.Run("custom logger but debug logging off", func(t *testing.T) {
		client, err = goobs.New("localhost:4545", goobs.WithPassword("hello"), goobs.WithLogger(customLogger{}), goobs.WithDebug(false))
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%T", client.Log), "*goobs.discard")
	})
}

func Test_goobs_client_errors(t *testing.T) {
	var client *goobs.Client
	var err error

	t.Run("basic", func(t *testing.T) {
		client, err = goobs.New("localhost:4545", goobs.WithPassword("hello"))
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%T", client.Log), "*goobs.discard")
	})

	t.Run("default logger used when only debug", func(t *testing.T) {
		client, err = goobs.New("localhost:4545", goobs.WithPassword("hello"), goobs.WithDebug(true))
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%T", client.Log), "*log.Logger")
	})

	t.Run("custom logger implies debug logging", func(t *testing.T) {
		client, err = goobs.New("localhost:4545", goobs.WithPassword("hello"), goobs.WithLogger(customLogger{}))
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%T", client.Log), "e2e.customLogger")
	})

	t.Run("custom logger but debug logging off", func(t *testing.T) {
		client, err = goobs.New("localhost:4545", goobs.WithPassword("hello"), goobs.WithLogger(customLogger{}), goobs.WithDebug(false))
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%T", client.Log), "*goobs.discard")
	})
}
