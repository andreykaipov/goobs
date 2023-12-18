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
	params := general.NewCallVendorRequestParams().
		WithVendorName("obs-shutdown-plugin").
		WithRequestType("shutdown").
		WithRequestData(map[string]interface{}{
			"reason":      "cleaning up",
			"support_url": "https://github.com/norihiro/obs-shutdown-plugin/issues",
			"force":       true,
		})
	resp, err := client.General.CallVendorRequest(params)
	fmt.Println(resp, err)
}
