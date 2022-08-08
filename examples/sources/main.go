package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/requests/inputs"
)

func main() {
	client, err := goobs.New(
		"localhost:4444",
		goobs.WithPassword("goodpassword"), // optional
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", client)

	go func() {
		for event := range client.IncomingEvents {
			switch e := event.(type) {
			case *events.ExitStarted:
				log.Printf("Server closing; gracefully shutting down")
				log.Printf("Bye")
				//os.Exit(0)
			case error:
				log.Printf("Error: %s", e)
			default:
				log.Printf("Unhandled event: %#v", e)
			}
		}
	}()

	defer client.Disconnect()

	getStatsResp, err := client.General.GetStats()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", getStatsResp)

	versionResp, err := client.General.GetVersion()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", versionResp)

	//rand.Seed(time.Now().UnixNano())
	list, _ := client.Inputs.GetInputList(&inputs.GetInputListParams{
		InputKind: "as",
	})

	for _, _ = range list.Inputs {
		if _, err := client.Inputs.SetInputVolume(&inputs.SetInputVolumeParams{
			InputName:      "asd",
			InputVolumeMul: rand.Float64(),
		}); err != nil {
			panic(err)
		}
	}
	time.Sleep(30 * time.Second)

	//	if len(list.Inputs) == 0 {
	//		fmt.Println("No sources!")
	//		os.Exit(0)
	//	}

	fmt.Println("Test toggling the mute status of the first source...")

	//	name := list.Sources[0].Name
	//	resp, _ := client.Sources.GetVolume(&sources.GetVolumeParams{Source: name})
	//	fmt.Printf("%s is muted? %t\n", name, resp.Muted)
	//
	//	_, _ = client.Sources.ToggleMute(&sources.ToggleMuteParams{Source: name})
	//	resp, _ = client.Sources.GetVolume(&sources.GetVolumeParams{Source: name})
	//	fmt.Printf("%s is muted? %t\n", name, resp.Muted)
}
