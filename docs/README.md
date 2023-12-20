# usage docs

- [getting started](#getting-started)
- [connecting](#connecting)
- [making requests](#making-requests)
  * [example 1](#example-1)
  * [example 2](#example-2)
  * [accessing the raw response](#accessing-the-raw-response)
- [configuration](#configuration)

## getting started

To get started, make sure you're running the latest OBS version so the [WebSocket Server](https://github.com/obsproject/obs-websocket) is already included alongside it.

If you don't see settings for the server under the _Tools_ menu like in the following screenshot, you're likely on a very old version and it's highly recommended to upgrade.

![screenshot of the OBS tools menu showing _WebSocket Server Settings_](obs-tools-menu.png)

For older distributions, it's still possible to run the Websocket server as a plugin, but chances are the server's protocol will be v4 and unusable with this library.
Please see [obsproject/obs-websocket](https://github.com/obsproject/obs-websocket) for more details on that.
The last goobs release to support v4 was [v0.8.1](https://github.com/andreykaipov/goobs/releases/tag/v0.8.1), so the following documentation doesn't apply and you'll be in for a rough ride!

## connecting

From the server settings, confirm the server is enabled:

![screenshot of the _WebSocket Server Settings_ modal window](obs-websocket-server-settings.png)

And connect!

```go
        client, err := goobs.New("localhost:4455", goobs.WithPassword("bigmoney420!!"))
```

Note for Windows users: if your development environment is within WSL, but OBS is running on the Windows host, you'll have to point the address to the Windows host instead, i.e. `ip route show default | awk '{print $3}'`.

The `goobs.WithPassword(...)` is an option we can pass to `goobs.New(...)` to configure the client.
Check the [Go docs](https://pkg.go.dev/github.com/andreykaipov/goobs#Option) for all the other possible options.

## making requests

The [OBS websocket protocol](https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md) is pretty exhaustive, and it's recommended to look through the requests and events there to see if any spark any ideas.
To help figure out how to perform a certain request with this library, a mapping between requests in the protocol to corresponding goobs requests is available at [request-mapping.md](/docs/request-mapping.md).

This client library is in fact [generated from the same protocol documentation](/internal/generate/protocol/main.go), so using goobs should hopefully remain fairly consistent with what you see in the protocol with little surprises.
If something seems off though, please open up an issue!

Figuring out what requests to use is mostly heuristic, but the following examples should hopefully guide the way and provide the necessary techniques to automate more complex tasks within OBS.

### example 1

Consider the common task of reordering some sources within a scene and toggling their visibility.

To figure out what requests we'd have to use to accomplish this, it's first recommended to listen to the events from the server and simply log them:

```go
        client.Listen(func(event any) {
                fmt.Printf("got event: %T\n", event)
        })
```

When that's running, let's switch over to OBS, and manually reorder some sources and toggle the visibility of a few.
Here's an example snippet of the output when we do that:

```console
got event: *events.SceneItemEnableStateChanged
got event: *events.SceneItemEnableStateChanged
got event: *events.SceneItemListReindexed
got event: *events.SceneItemSelected
got event: *events.SceneItemSelected
got event: *events.SceneItemListReindexed
got event: *events.SceneItemEnableStateChanged
```

There's no documented mapping between requests and events, but they tend to be similarly named, so this output can help clue us into what request category to look into.
In this case, we should look at _Scene Items_.

Now refer to the [obs-websocket docs](https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#scene-items-requests) and our [goobs mapping](/docs/request-mapping.md#scene-items) to browse through the list of available requests to narrow it down further.
In our output snippet above, the scene item events of interest are `SceneItemEnableStateChanged` and `SceneItemListReindexed`, and thankfully there's two aptly named requests - `SetSceneItemEnabled` and `SetSceneItemIndex`.

Consider `SetSceneItemEnabled`.
It's [documentation](https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#setsceneitemenabled) says the request expects the scene name, the scene item ID, and the new enabled state of the scene item.
But what's the ID of a scene item?

We might look through the request list again to find the `GetSceneItemId` and `GetSceneItemList` requests.
These will work, and in fact we'll make use of `GetSceneItemList` later, but for the sake of example, consider the following approach.

Now that we know the events of interest, instead of just printing out their type, we can inspect them more in depth by using type assertion in our `client.Listen` callback and printing out the details:

```go
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
```

After reordering and toggling some scene items in OBS, the output might look as follows:

```console
reindexed: Scene
  &{SceneItemID:47 SceneItemIndex:0}
  &{SceneItemID:54 SceneItemIndex:1}
  &{SceneItemID:53 SceneItemIndex:2}
visibility:
  &{SceneItemEnabled:false SceneItemId:47 SceneName:Scene}
unhandled: *events.SceneItemSelected
reindexed: Scene
  &{SceneItemID:53 SceneItemIndex:0}
  &{SceneItemID:47 SceneItemIndex:1}
  &{SceneItemID:54 SceneItemIndex:2}
visibility:
  &{SceneItemEnabled:true SceneItemId:47 SceneName:Scene}
```

Alrighty, so now we know what the IDs of our sources look like, and while we could hardcode them into the parameters of whatever requests we want to use, let's go for a friendlier approach.
In this case, looping over our scene items and finding the source based on a given name sounds like a reasonable solution.

It might look something like this:

```go
        sceneName := os.Getenv("SCENE_NAME")
        sourceName := os.Getenv("SOURCE_NAME")

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
```

And the source for `setSourceVisibility`:

```go
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
```

Note the use of the builder pattern to define the params for `GetSceneItemList`, while `SetSceneItemEnabled` created the params struct directly.
Use whichever you like best!

We can also send our `client.Listen` block to the background in a goroutine at the same time so we can keep an eye on the events.
The final code is available in the [examples under docs-demo](/_examples/docs-demo/main.go):

```console
❯ OBS_HOST=$WINHOST:4455 SCENE_NAME=Scene SOURCE_NAME=Image go run _examples/test/main.go
visibility:
  &{SceneItemEnabled:false SceneItemId:53 SceneName:Scene}
visibility:
  &{SceneItemEnabled:false SceneItemId:47 SceneName:Scene}
visibility:
  &{SceneItemEnabled:false SceneItemId:54 SceneName:Scene}
visibility:
  &{SceneItemEnabled:true SceneItemId:54 SceneName:Scene}
```

### example 2

For this example, consider the [CreateInput](https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#createinput) request under the _Inputs_ category.
The following will add a new video capture input as a source into our specified scene:

```go
        params := inputs.NewCreateInputParams().
                WithSceneName("my scene").
                WithInputName("new browser").
                WithInputKind("dshow_input")
        resp, err := client.Inputs.CreateInput(params)
```

But wait a minute... how did you know to use `dshow_input`?

Same idea as in the first example!

For example, we could use `client.Inputs.GetInputList()` to iterate over all the inputs to view the various "input kinds" of our existing inputs to narrow it down:

```go
        resp, _ := client.Inputs.GetInputList()
        for _, i := range resp.Inputs {
                fmt.Printf("%+v\n", i)
        }
```

We can also subscribe to the server events, manually interact with OBS, and see what gets printed out, hoping it has the info we're looking for.

For the sake of example, instead of using `client.Listen()`, we can also read from the `client.IncomingEvents` channel directly:

```go
        for _, event := range c.IncomingEvents {
                fmt.Printf("got event: %#v\n", event)
        }
```

The idea is the same though.
If we create a video capture input, we'll find the following event in our logs:

```
got event: &events.InputCreated{DefaultInputSettings:map[string]interface {}{"active":true, "audio_output_mode":0, "autorotation":true, "color_range":"default", "color_space":"default", "frame_interval":-1, "hw_decode":false, "res_type":0, "video_format":0}, InputKind:"dshow_input", InputName:"Video Capture Device 4", InputSettings:map[string]interface {}{}, UnversionedInputKind:"dshow_input"}
```

And in that blob of JSON, we'll find `InputKind:"dshow_input"`!

### accessing the raw response

If you find the generated API to be incorrect, and the response for a request is missing documented fields in the Go struct, accessing the "raw" response from the server is possible, so that you can perform the necessary JSON unmarshalling yourself:

```go
        resp, _ := client.Inputs.GetInputList()
        data := map[string]any{}
        if err := json.Unmarshal(resp.GetRaw(), &data); err != nil {
                panic(err)
        }
        fmt.Println(data["inputs"])
```

If you find yourself having to use this though, please open up an issue!

## configuration

The following environment variables can be set to configure goobs:

- `GOOBS_LOG` can be set to one of `trace`, `debug`, `info`, or `error` to better understand what our client is doing under the hood.

- `GOOBS_PROFILE` can be set to enable profiling.
  For example, the following will help us find unreleased memory:

  ```console
  ❯ GOOBS_PROFILE=memprofile=mem.out OBS_PORT=4455 go test -v -run=profile client_test.go
  ❯ go tool pprof -top -sample_index=inuse_space mem.out
  ```

  Set `GOOBS_PROFILE=help` to see all the other available options.
