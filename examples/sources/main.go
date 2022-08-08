package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/inputs"
)

func blue(f string, v ...any)  { log.Printf("\033[34m"+f+"\033[0m", v...) }
func green(f string, v ...any) { log.Printf("\033[32m"+f+"\033[0m", v...) }

func main() {
	client, _ := goobs.New("localhost:4444", goobs.WithPassword("goodpassword"))
	defer client.Disconnect()

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

		resp, _ := client.Inputs.GetInputVolume(&inputs.GetInputVolumeParams{InputName: name})
		green("%s's volume is now %f", name, resp.InputVolumeDb)
	}
}
