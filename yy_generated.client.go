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
	c.General = general.NewClient(general.WithConn(c.conn))
	c.Profiles = profiles.NewClient(profiles.WithConn(c.conn))
	c.Recording = recording.NewClient(recording.WithConn(c.conn))
	c.ReplayBuffer = replaybuffer.NewClient(replaybuffer.WithConn(c.conn))
	c.SceneCollections = scenecollections.NewClient(scenecollections.WithConn(c.conn))
	c.SceneItems = sceneitems.NewClient(sceneitems.WithConn(c.conn))
	c.Scenes = scenes.NewClient(scenes.WithConn(c.conn))
	c.Sources = sources.NewClient(sources.WithConn(c.conn))
	c.Streaming = streaming.NewClient(streaming.WithConn(c.conn))
	c.StudioMode = studiomode.NewClient(studiomode.WithConn(c.conn))
	c.Transitions = transitions.NewClient(transitions.WithConn(c.conn))
}
