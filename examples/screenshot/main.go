package main

import (
	"encoding/base64"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/sources"
)

func main() {
	client, err := goobs.New("localhost:4455", goobs.WithPassword("goodpassword"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	params := sources.NewGetSourceScreenshotParams().
		WithSourceName("Webcam").
		WithImageCompressionQuality(-1).
		WithImageFormat("png")
	screenshot, err := client.Sources.GetSourceScreenshot(params)
	if err != nil {
		log.Fatal(err)
	}

	// remove the `data:image/png;base64,...` prefix
	data := screenshot.ImageData[strings.IndexByte(screenshot.ImageData, ',')+1:]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	image, _ := png.Decode(reader)
	f, _ := os.Create("example.png")
	_ = png.Encode(f, image)
}
