package main

import (
	"fmt"
	"os"
	"time"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/requests/sceneitems"
)

var client *goobs.Client
var err error

func main() {
	host := os.Getenv("OBS_HOST")
	sceneName := os.Getenv("SCENE_NAME")
	sourceName := os.Getenv("SOURCE_NAME")

	client, err = goobs.New(host, goobs.WithPassword("goodpassword"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	go listen()

	params := sceneitems.NewGetSceneItemListParams().WithSceneName(sceneName)
	sceneList, err := client.SceneItems.GetSceneItemList(params)
	if err != nil {
		panic(err)
	}

	// find the ID of our source, while hiding all others
	var sourceID int
	for _, item := range sceneList.SceneItems {
		if item.SourceName == sourceName {
			sourceID = item.SceneItemID
		}

		setSourceVisibility(sceneName, item.SceneItemID, false)
	}

	// then show our source
	setSourceVisibility(sceneName, sourceID, true)

	// wait for events to come through
	time.Sleep(1 * time.Second)
}

func listen() {
	client.Listen(func(event any) {
		switch e := event.(type) {
		case *events.SceneItemListReindexed:
			fmt.Printf("reindexed: %v\n", e.SceneName)
			for _, item := range e.SceneItems {
				fmt.Printf("  %+v\n", item)
			}
		case *events.SceneItemEnableStateChanged:
			fmt.Printf("visibility:\n")
			fmt.Printf("  %+v\n", e)
		default:
			fmt.Printf("unhandled: %T\n", event)
		}
	})
}

func setSourceVisibility(scene string, sourceID int, visible bool) {
	params := &sceneitems.SetSceneItemEnabledParams{
		SceneName:        &scene,
		SceneItemId:      &sourceID,
		SceneItemEnabled: &visible,
	}
	_, err := client.SceneItems.SetSceneItemEnabled(params)
	if err != nil {
		panic(err)
	}
}
