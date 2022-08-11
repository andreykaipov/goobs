package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/requests/inputs"
)

func green(f string, v ...any) { log.Printf("\033[32m"+f+"\033[0m", v...) }
func blue(f string, v ...any)  { log.Printf("\033[34m"+f+"\033[0m", v...) }

func main() {
	client, err := goobs.New("localhost:4455", goobs.WithPassword("goodpassword"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	go client.Listen(func(event any) {
		switch e := event.(type) {
		case *events.InputVolumeChanged:
			green("%s's volume is now %f", e.InputName, e.InputVolumeDb)
		case *events.ExitStarted:
			green("Gracefully shutting down")
			green("Bye")
			os.Exit(0)
		default:
			blue("Unhandled: %#v", e)
		}
	})

	rand.Seed(time.Now().UnixNano())

	list, _ := client.Inputs.GetInputList()

	for _, v := range list.Inputs {
		name := v.InputName

		if _, err := client.Inputs.SetInputVolume(&inputs.SetInputVolumeParams{
			InputName:      name,
			InputVolumeMul: rand.Float64(),
		}); err != nil {
			blue("%s doesn't support audio", name)
			continue
		}
	}

	// play around in OBS and maybe close it - observe the events it sends
	//
	time.Sleep(30 * time.Second)
}
