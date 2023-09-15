package main

import (
	"fmt"
	"log"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/general"
)

func main() {
	client, err := goobs.New("localhost:4455", goobs.WithPassword("goodpassword"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	// this assumes you have the obs-shutdown-plugin installed
	// see https://github.com/norihiro/obs-shutdown-plugin
	//
	resp, err := client.General.CallVendorRequest(&general.CallVendorRequestParams{
		VendorName:  "obs-shutdown-plugin",
		RequestType: "shutdown",
		RequestData: map[string]interface{}{
			"reason":      "cleaning up",
			"support_url": "https://github.com/norihiro/obs-shutdown-plugin/issues",
			"force":       true,
		},
	})
	fmt.Println(resp, err)
}
