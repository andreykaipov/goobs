// This file has been automatically generated. Don't edit it.

package goobs

import (
	general "github.com/andreykaipov/goobs/api/requests/general"
	mediacontrol "github.com/andreykaipov/goobs/api/requests/media_control"
	outputs "github.com/andreykaipov/goobs/api/requests/outputs"
	profiles "github.com/andreykaipov/goobs/api/requests/profiles"
	recording "github.com/andreykaipov/goobs/api/requests/recording"
	replaybuffer "github.com/andreykaipov/goobs/api/requests/replay_buffer"
	scenecollections "github.com/andreykaipov/goobs/api/requests/scene_collections"
	sceneitems "github.com/andreykaipov/goobs/api/requests/scene_items"
	scenes "github.com/andreykaipov/goobs/api/requests/scenes"
	sources "github.com/andreykaipov/goobs/api/requests/sources"
	streaming "github.com/andreykaipov/goobs/api/requests/streaming"
	studiomode "github.com/andreykaipov/goobs/api/requests/studio_mode"
	transitions "github.com/andreykaipov/goobs/api/requests/transitions"
	virtualcam "github.com/andreykaipov/goobs/api/requests/virtual_cam"
)

type subclients struct {
	General          *general.Client
	MediaControl     *mediacontrol.Client
	Outputs          *outputs.Client
	Profiles         *profiles.Client
	Recording        *recording.Client
	ReplayBuffer     *replaybuffer.Client
	SceneCollections *scenecollections.Client
	SceneItems       *sceneitems.Client
	Scenes           *scenes.Client
	Sources          *sources.Client
	Streaming        *streaming.Client
	StudioMode       *studiomode.Client
	Transitions      *transitions.Client
	VirtualCam       *virtualcam.Client
}

func setClients(c *Client) {
	c.General = &general.Client{Client: c.Client}
	c.MediaControl = &mediacontrol.Client{Client: c.Client}
	c.Outputs = &outputs.Client{Client: c.Client}
	c.Profiles = &profiles.Client{Client: c.Client}
	c.Recording = &recording.Client{Client: c.Client}
	c.ReplayBuffer = &replaybuffer.Client{Client: c.Client}
	c.SceneCollections = &scenecollections.Client{Client: c.Client}
	c.SceneItems = &sceneitems.Client{Client: c.Client}
	c.Scenes = &scenes.Client{Client: c.Client}
	c.Sources = &sources.Client{Client: c.Client}
	c.Streaming = &streaming.Client{Client: c.Client}
	c.StudioMode = &studiomode.Client{Client: c.Client}
	c.Transitions = &transitions.Client{Client: c.Client}
	c.VirtualCam = &virtualcam.Client{Client: c.Client}
}
