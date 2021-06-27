package e2e

import (
	"testing"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/stretchr/testify/assert"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_goobs_e2e(t *testing.T) {
	client, err := goobs.New("localhost:4545", goobs.WithPassword("hello"))
	assert.NoError(t, err)

	sceneName := "goobs test scene"

	t.Run("Scenes", func(t *testing.T) {
		resp, err := client.Scenes.GetSceneList()
		assert.NoError(t, err)
		list := resp.Scenes

		for _, s := range list {
			if s.Name == sceneName {
				// Prior OBS instance probably hasn't been closed
				t.Logf("Scene list already contains e2e scene %q.", sceneName)
			}
		}

		_, err = client.Scenes.CreateScene(&scenes.CreateSceneParams{SceneName: sceneName})
		assert.NoError(t, err)
		_, err = client.Scenes.CreateScene(&scenes.CreateSceneParams{SceneName: sceneName})
		assert.Error(t, err)

		resp, err = client.Scenes.GetSceneList()
		assert.NoError(t, err)
		newList := resp.Scenes
		assert.Len(t, newList, len(list)+1)
	})

	// 	resp, _ := client.Sources.GetTextGDIPlusProperties(&sources.GetTextGDIPlusPropertiesParams{Source: "hello"})
	// 	//fmt.Printf("Text is %s in font face %s\n", resp.Text, resp.Font.Face)
	//
	// 	//params := &sources.SetTextGDIPlusPropertiesParams{}
	// 	//params.Source = "hello"
	// 	//params.Text = "abc"
	// 	//params.Font.Face = "Arial Black"
	// 	//params.Font.Size = 20
	//
	// 	//	resp, _ = client.Sources.GetTextGDIPlusProperties(&sources.GetTextGDIPlusPropertiesParams{Source: "hello"})
	// 	//	fmt.Printf("%#v\n", resp.Font)
	//
	// 	resp1, err := client.Sources.SetTextGDIPlusProperties(&sources.SetTextGDIPlusPropertiesParams{
	// 		Source: "hello",
	// 		Text:   "abc",
	// 		Font: &typedefs.Font{
	// 			Face: "Arial Black",
	// 			Size: 11,
	// 		},
	// 	})
	// 	fmt.Printf("%#v\n", resp1)
	// 	fmt.Printf("%#v\n", err)
	//
	// 	resp, _ = client.Sources.GetTextGDIPlusProperties(&sources.GetTextGDIPlusPropertiesParams{Source: "hello"})
	// 	fmt.Printf("%#v\n", resp.Font)
	// 	fmt.Printf("Now text is %s in font face %s\n", resp.Text, resp.Font.Face)

	//resp2, err := client.SceneItems.GetSceneItemProperties(&sceneitems.GetSceneItemPropertiesParams{
	//	Item:      &typedefs.Item{Name: "Chat"},
	//	SceneName: "Scene",
	//})

	//	resp2, err := client.MediaControl.GetMediaDuration(&mediacontrol.GetMediaDurationParams{})
	//	fmt.Printf("%#v\n", resp2)
	//	fmt.Printf("%#v\n", err)
}
