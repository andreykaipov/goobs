package main

import (
	"fmt"
	"log"
	"time"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/inputs"
	"github.com/andreykaipov/goobs/api/requests/sceneitems"
)

var client *goobs.Client

func list() {
	fmt.Println("listing scene items...")
	items, err := client.SceneItems.GetSceneItemList(&sceneitems.GetSceneItemListParams{SceneName: "Scene"})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range items.SceneItems {
		fmt.Printf("index=%d, id=%d, type=%s, name=%s\n", v.SceneItemIndex, v.SceneItemID, v.SourceType, v.SourceName)
	}
}

func create(scene, name, kind string) float64 {
	resp, err := client.Inputs.CreateInput(&inputs.CreateInputParams{SceneName: scene, InputName: name, InputKind: kind})
	if err != nil {
		log.Fatal(err)
	}
	return resp.SceneItemId
}

func main() {
	var err error
	client, err = goobs.New("localhost:4455", goobs.WithPassword("goodpassword"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	create("Scene", "test1", "browser_source")
	create("Scene", "test2", "game_capture")
	create("Scene", "test3", "dshow_input")
	create("Scene", "test4", "image_source")

	list()

	fmt.Println("sleeping for 3 seconds...")
	time.Sleep(3 * time.Second)

	fmt.Println("reversing the order of each scene item in the scene")
	items, _ := client.SceneItems.GetSceneItemList(&sceneitems.GetSceneItemListParams{SceneName: "Scene"})
	for _, v := range items.SceneItems {
		_, _ = client.SceneItems.SetSceneItemIndex(&sceneitems.SetSceneItemIndexParams{
			SceneName:      "Scene",
			SceneItemId:    float64(v.SceneItemID),
			SceneItemIndex: float64(len(items.SceneItems) - v.SceneItemIndex - 1),
		})
	}

	list()
}
