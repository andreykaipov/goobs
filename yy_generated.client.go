// This file has been automatically generated. Don't edit it.

package goobs

import (
	general "github.com/andreykaipov/goobs/api/requests/general"
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
)

type subclients struct {
	General          *general.Client
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
}

func setClients(c *Client) {
	c.General = general.NewClient(general.WithConn(c.requestsConn))
	c.Profiles = profiles.NewClient(profiles.WithConn(c.requestsConn))
	c.Recording = recording.NewClient(recording.WithConn(c.requestsConn))
	c.ReplayBuffer = replaybuffer.NewClient(replaybuffer.WithConn(c.requestsConn))
	c.SceneCollections = scenecollections.NewClient(scenecollections.WithConn(c.requestsConn))
	c.SceneItems = sceneitems.NewClient(sceneitems.WithConn(c.requestsConn))
	c.Scenes = scenes.NewClient(scenes.WithConn(c.requestsConn))
	c.Sources = sources.NewClient(sources.WithConn(c.requestsConn))
	c.Streaming = streaming.NewClient(streaming.WithConn(c.requestsConn))
	c.StudioMode = studiomode.NewClient(studiomode.WithConn(c.requestsConn))
	c.Transitions = transitions.NewClient(transitions.WithConn(c.requestsConn))
}
