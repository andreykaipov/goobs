package goobs_test

import (
	"net"
	"net/http"
	"os"
	"sync"
	"testing"

	goobs "github.com/andreykaipov/goobs"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func Test_client(t *testing.T) {
	var err error
	_, err = goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("wrongpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.Error(t, err)
	assert.IsType(t, &websocket.CloseError{}, err)
	assert.Equal(t, err.(*websocket.CloseError).Code, 4009)
	_, err = goobs.New(
		"localhost:42069",
		goobs.WithPassword("wrongpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.Error(t, err)
	assert.IsType(t, &net.OpError{}, err)
}

func Test_multi_goroutine(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := client.Scenes.GetSceneList()
			assert.NoError(t, err)
		}()
	}
	wg.Wait()
}
